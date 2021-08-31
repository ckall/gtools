package push

import (
	"errors"
	"gck/push/ding"
)

type DingTalk interface {
	Markdown(title, context string, opt ...ding.AtOption) *dingMarkdown
	Text(context string, opt ...ding.AtOption) *dingText
}

type DingMarkdown interface {
	Send() error
}

type DingText interface {
	Send() error
}

var (
	//初始化ding模块
	initDingTalk = make(map[string]*ding.DingTalk)
)

type dingTalk struct {
	name string
}

//公共消息体
type dingPublicMessageBody struct {
	name    string
	context string
	opt     []ding.AtOption
}

//MarkDown消息类型
type dingMarkdown struct {
	title string
	dingPublicMessageBody
}

//Text消息类型
type dingText = dingPublicMessageBody

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

//ding模块
func Ding(name string) *dingTalk {
	return &dingTalk{
		name: name,
	}
}

//发送Markdown消息
func (ding *dingTalk) Markdown(title, context string, opt ...ding.AtOption) *dingMarkdown {
	initMarkDown := &dingMarkdown{
		title: title,
	}
	initMarkDown.name = ding.name
	initMarkDown.context = context
	initMarkDown.opt = opt
	return initMarkDown
}

//发送Text消息
func (ding *dingTalk) Text(context string, opt ...ding.AtOption) *dingText {
	initText := &dingText{
		name: ding.name,
	}
	initText.context = context
	initText.opt = opt
	return initText
}

//发送
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
