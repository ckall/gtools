# gck 

A large collection of golang general modules, integrated log system, warning push system, performance monitoring golang 通用模块大集合，集成日志系统，警告推送系统，性能监控中心


日志系统  
```
  日志分割
  格式定义
```
推送系统
```
  信息框警告：
     钉钉   
  邮箱警告
     待完成    
  电话警告 
     待完成      
```
### 链路跟踪:
```
   * zipkin 
```

## 日志模块

```

```

## 钉钉的使用


### dingTalk
```go
   //支持原生语句
   	//context.AddText("###### 10点20分发布 [天气](http://www.thinkpage.cn/)")
   	context := ding.NewConText()
    //如果没有找到支持的文本格式，可以用到官方的格式
   	context.AddText("# 杭州天气")
    //如果对字体有要求的情况
   	context.AddText(ding.AddH3("9度，西北风1级，空气良89，相对温度73%"))
   	//可以用到官方的格式
   	context.AddKeyValue("#### 【测试】:", map[string]interface{}{"测": "试"})
    //支持随意的格式
    context.AddKeyValue(ding.AddH4("【测试】:"), map[string]interface{}{"测": "试"})
   	//添加图片
    context.AddImage("https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png")
   	//添加图片
    context.AddImage("https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png")
   	//添加文字并且还要颜色
    context.AddTextUrl(ding.AddH1("10点20分发布: %s 和 %s "), map[string]string{ding.AddGreen("天气"): "http://www.thinkpage.cn/", ding.AddRed("天气11"): "http://www.baidu.com/"})
   	//添加文本
    context.AddText(ding.AddH6("杭 %s 和 %s"), ding.AddRed("hiehie"), ding.AddBlue("hiehie"))
   	context.AddText("杭 %s 和 %s", ding.AddGreen("hiehie"), ding.AddGold("hiehie"))
   	//发送
    //name 表示要发送到那个机器人里面
    ding.WithAtMobiles([]string{"1731122967*"} //@人员
    ding.WithAtAll() //@全体
    err := push.Ding(name).
   		Markdown(title, context, ding.WithAtMobiles([]string{"1731122967*"})).
   		Send()
   	if err != nil {
   		t.Errorf("发送失败: %s", err.Error())
   	}
```  

## 性能监控

##### 访问http://HOST:PORT/debug/pprof

##### 内存火焰图本地运行go tool pprof -http=:8081 http://HOST:PORT/debug/pprof/heap

##### CPU火焰图本地运行go tool pprof -http=:8081 http://HOST:PORT/debug/pprof/profile?seconds=10

##### 实时监控访问http://HOST:PORT/debug/charts