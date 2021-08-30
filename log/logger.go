package log

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"strings"
	"time"
)

const (
	DefaultEnvType            = "dev" //默认的开发环境
	DefaultFileCount    uint  = 0     //0:为不限制
	DefaultFileSize     int64 = 32 * 1024 * 1024
	DefaultFilePath           = "./glogs/%Y-%m-%d.log" //日志路径
	DefaultTemplateJoin       = " "                    //日志分割符 msg:"信息1 信息2"
	//错误级别
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

var (
	gck = &logConfig{
		onlyFileSize: DefaultFileSize,
		envType:      DefaultEnvType,
		logPath:      DefaultFilePath,
		fileCount:    DefaultFileCount,
		TemplateJoin: DefaultTemplateJoin,
	}
	zaplog *zap.Logger
)

/**
 * @auth: kuncheng
 * @Date: 2021/8/28
 */
type logConfig struct {
	onlyFileSize int64         //单个文件大小(大小限制)
	envType      string        //环境类型
	fileCount    uint          //文件总数
	logPath      string        //文件路径
	logType      string        //输出形式（日志，工作台，日志 工作台（默认））
	rotationTime time.Duration // 日志分割的时间
	maxAge       time.Duration // 日志最大保留的天数
	TemplateJoin string
	logFields    map[string]interface{} //日志扩展字段
}

//钉钉配置
type dingDingConfig struct {
	token string
	host  string
}

//配置回调
type LogOptionsFunc func(*logConfig)

func InitLog(options ...LogOptionsFunc) {
	for _, option := range options {
		option(gck)
	}
	initZapLog(gck)
}

//初始化扩展
func initZapLog(lc *logConfig) {
	//格式初始化
	encoderConfig := zap.NewProductionEncoderConfig()
	var defaultLogLevel = zap.NewAtomicLevel()
	encoderConfig.EncodeTime = timeEncoder
	logPath := fmt.Sprintf("%s", lc.logPath)
	w := zapcore.AddSync(GetWriter(
		logPath,
		WithRotationSize(lc.onlyFileSize),
		WithRotationCount(lc.fileCount),
		WithRotationTime(lc.rotationTime),
		WithMaxAge(lc.maxAge),
	))
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		defaultLogLevel,
	)
	//添加日志字段
	fields := []zap.Field{
		zap.String("env_type", lc.envType),
	}
	for key, field := range lc.logFields {
		fields = append(fields, zap.Reflect(key, field))
	}
	// 构建日志
	zaplog = zap.New(core, zap.Fields(fields...))
}

// 按天切割按大小切割
// filename 文件名
// rotationSize 每个文件的大小
// maxAge 文件最大保留天数
// rotationCount 最大保留文件个数
// rotationTime 设置文件分割时间
// rotationCount 设置保留的最大文件数量
func GetWriter(filename string, options ...Option) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 stream-2021-5-20.log
	// demo.log是指向最新日志的连接
	// 保存7天内的日志，每1小时(整点)分割一第二天志
	hook, err := New(
		filename,
		options...,
	)

	if err != nil {
		panic(err)
	}
	return hook
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	var layout = "2006-01-02 15:04:05"
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}

//数据处理
//
func handleArgs(templates []interface{}) (msg string) {
	var elems []string
	for _, template := range templates {
		var str string
		var ok bool
		if str, ok = template.(string); !ok {
			byteStr, err := json.Marshal(template)
			if err != nil {
				panic("数据无法解析")
				return
			}
			str = string(byteStr)
		}
		elems = append(elems, str)
	}
	msg = strings.Join(elems, gck.TemplateJoin)
	return
}

//设置单个文件大小(大小限制)
func SetOnlyFileSize(onlyFileSize int64) LogOptionsFunc {
	return func(c *logConfig) {
		if onlyFileSize != 0 {
			c.onlyFileSize = onlyFileSize
		}
	}
}

// 设置保留的最大文件数量、没有默认值(表示不限制)
func SetRotationCount(n time.Duration) LogOptionsFunc {
	return func(c *logConfig) {
		if n != 0 {
			c.rotationTime = n
		}
	}
}

//设置环境类型
func SetEnvType(envType string) LogOptionsFunc {
	return func(c *logConfig) {
		if envType != "" {
			c.envType = envType
		}
	}
}

//日志最大保留的天数
func SetMaxAge(maxAge time.Duration) LogOptionsFunc {
	return func(c *logConfig) {
		if maxAge != 0 {
			c.maxAge = maxAge
		}
	}
}

//设置文件总数
func SetFileCount(fileCount uint) LogOptionsFunc {
	return func(c *logConfig) {
		if fileCount != 0 {
			c.fileCount = fileCount
		}
	}
}

//设置文件路径
func SetFilePath(logPath string) LogOptionsFunc {
	return func(c *logConfig) {
		if logPath != "" {
			c.logPath = logPath
		}
	}
}

//日志扩展字段
func SetLogField(key string, value interface{}) LogOptionsFunc {
	return func(c *logConfig) {
		c.logFields[key] = value
	}
}

//日志扩展字段
//不会追加 采用覆盖
func SetLogFields(field map[string]interface{}) LogOptionsFunc {
	return func(c *logConfig) {
		c.logFields = field
	}
}

//日志信息
func Info(template ...interface{}) {
	var message = handleArgs(template)
	writer(context.Background(), LevelInfo, message)
}

//日志信息
func Warn(template ...interface{}) {
	var message = handleArgs(template)
	writer(context.Background(), LevelWarn, message)
}

//日志信息
func Error(template ...interface{}) {
	var message = handleArgs(template)
	writer(context.Background(), LevelError, message)
}

func writer(ctx context.Context, levelType string, message string, field ...zap.Field) {
	switch levelType {
	case LevelInfo:
		zaplog.Info(message, field...)
	case LevelWarn:
		zaplog.Warn(message, field...)
	case LevelError:
		zaplog.Error(message, field...)
	default:
		panic("未知的日志类型级别")
	}
}
