package push_test

import (
	"gck/push"
	"gck/push/ding"
	"testing"
)

var name = "test"
var token = "5a3eaeb2d3f340da2fca857a58654eef325224b4d1f68611bd792da31b452215"
var title = "test"

func init() {
	push.InitKeyWordDing(
		name,
		token,
	)
}

func TestDingTalk_Markdown(t *testing.T) {
	//支持原生语句
	//context.AddText("###### 10点20分发布 [天气](http://www.thinkpage.cn/)")
	context := ding.NewConText()
	//context.AddText("# 杭州天气")
	//context.AddText("### 9度，西北风1级，空气良89，相对温度73%")
	//context.AddKeyValue("#### 【测试】:", map[string]interface{}{"测": "试"})
	//context.AddImage("https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png")
	context.AddTextUrl(ding.AddH1("10点20分发布: %s 和 %s "), map[string]string{ding.AddGreen("天气"): "http://www.thinkpage.cn/", ding.AddRed("天气11"): "http://www.baidu.com/"})
	context.AddText(ding.AddH6("杭 %s 和 %s"), ding.AddRed("hiehie"), ding.AddBlue("hiehie"))
	//context.AddText("杭 %s 和 %s", ding.AddGreen("hiehie"), ding.AddGold("hiehie"))
	err := push.Ding(name).
		Markdown(title, context, ding.WithAtMobiles([]string{"1731122967*"})).
		Send()
	if err != nil {
		t.Errorf("发送失败: %s", err.Error())
	}
}

func TestDingTalk_Text(t *testing.T) {
	err := push.Ding(name).Text("测试文本消息", ding.WithAtAll()).Send()
	if err != nil {
		t.Errorf("发送失败: %s", err.Error())
	}
}
