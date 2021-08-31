# 推送系统
<h3> 如何使用??? </h3>

### dingTalk
```go
    func init() {
	    log.InitLog()
        push.InitKeyWordDing("test", "5a3eaeb2d3f340da2fca857a58654eef325224b4d1f68611bd792da31b452215", "test")
    }
    func main() {
        err := push.Ding("test").Markdown("模块测试","text").Send()
            if err != nil {
                log.Error("模块发送失败:", err.Error())
                return
        }	
    }
```  

### zipkin
