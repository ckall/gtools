package push

import (
	"errors"
	"gck/push/ding"
)

var (
	initDingTalk = make(map[string]*ding.DingTalk)
)

type dingTalk struct {
	name string
}

type dingPublicMessageBody struct {
	context string
	opt     []ding.AtOption
}

//
type dingMarkdown struct {
	name  string
	title string
	dingPublicMessageBody
}

//
type dingText struct {
	name string
	dingPublicMessageBody
}

/**
 * @auth: kuncheng
 * @Date: 2021/8/30
 */
//关键词初始化
//tokens 可以发送给多个群
//key 每个群的多个关键字验证
func InitKeyWordDing(name, tokens string, title string) {
	initDingTalk[name] = ding.InitDingTalk(tokens, title)
}

//签名规则
func InitSecretDing(name, token string, secret string) {
	initDingTalk[name] = ding.InitDingTalkWithSecret(token, secret)
}

func Ding(name string) *dingTalk {
	return &dingTalk{
		name: name,
	}
}

//发送Markdown消息
func (ding *dingTalk) SendMarkdown(title, context string, opt ...ding.AtOption) *dingMarkdown {
	initMarkDown := &dingMarkdown{
		name:  ding.name,
		title: title,
	}
	initMarkDown.context = context
	initMarkDown.opt = opt
	return initMarkDown
}

//发送Text消息
func (ding *dingTalk) SendText(context string, opt ...ding.AtOption) *dingText {
	initText := &dingText{
		name: ding.name,
	}
	initText.context = context
	initText.opt = opt
	return initText
}

func (ding *dingText) Send() error {
	if talk, ok := initDingTalk[ding.name]; ok {
		return talk.SendTextMessage(ding.context, ding.opt...)
	}
	return errors.New("类型错误")
}

func (ding *dingMarkdown) Send() error {
	if talk, ok := initDingTalk[ding.name]; ok {
		return talk.SendMarkDownMessage(ding.title, ding.context, ding.opt...)
	}
	return errors.New("类型错误")
}
