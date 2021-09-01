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
	err := push.Ding(name).
		Markdown(title, "a", ding.WithAtMobiles([]string{"17311229673"})).
		Send()
	if err != nil {
		t.Errorf("发送失败: %s", err.Error())
	}
}

func TestDingTalk_Text(t *testing.T) {
	err := push.Ding(name).Text(title, ding.WithAtAll()).Send()
	if err != nil {
		t.Errorf("发送失败: %s", err.Error())
	}
}
