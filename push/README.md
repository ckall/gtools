# 推送系统
<h3> 如何使用??? </h3>

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

### zipkin
