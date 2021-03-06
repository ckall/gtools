package log_test

import (
	"github.com/ckall/gtools/log"
	"testing"
)

func init() {
	log.InitLog(
		log.SetLogField("name1", "无敌大帅比"),
		log.SetEnvType("test"),
		log.SetOnlyFileSize(1),
		log.SetFilePath("/var/log/app/%Y-%m-%d %H.log"),
	)
}

func TestInfo(t *testing.T) {
	log.Info("这个是", []string{"一个"}, []rune("简单"), "的", map[string]interface{}{
		"测": "试",
	})
}

func TestError(t *testing.T) {
	log.Error("这个是", []string{"一个"}, []rune("简单"), "的", map[string]interface{}{
		"测": "试",
	})
}

func TestWarn(t *testing.T) {
	log.Warn("这个是", []string{"一个"}, []rune("简单"), "的", map[string]interface{}{
		"测": "试",
	})
}
