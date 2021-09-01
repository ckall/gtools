package push

import (
	"errors"
	"gck/push/ding"
)

type DingTalk interface {
	Markdown(title, context string, opt ...ding.AtOption) Markdown

	Text(context string, opt ...ding.AtOption) Text

	LinkMessage(title string, text string, picUrl string, msgUrl string) LinkMessage
}

type name struct {
}

type LinkMessage interface {
	Send() error
}

type Markdown interface {
	Send() error
}

type Text interface {
	Send() error
}

var (
	//初始化ding模块
	initDingTalk = make(map[string]*ding.DingTalk)
)

type dingTalk struct {
	name string
}

type uRLLink struct {
	dingTalk
	title, text, picUrl, msgUrl string
}

//公共消息体
type publicMessageBody struct {
	dingTalk
	context string
	opt     []ding.AtOption
}

//MarkDown消息类型
type markdown struct {
	title string
	publicMessageBody
}

//Text消息类型
type text = publicMessageBody

/**
 * @auth: kuncheng
 * @Date: 2021/8/30
 */
//关键词初始化
//tokens 可以发送给多个群
//key 每个群的多个关键字验证
func InitKeyWordDing(name, token string) {
	initDingTalk[name] = ding.InitDingTalk(token)
}

//签名规则
func InitSecretDing(name, token string, secret string) {
	initDingTalk[name] = ding.InitDingTalkWithSecret(token, secret)
}

//ding模块
func Ding(name string) DingTalk {
	return &dingTalk{
		name: name,
	}
}

//发送Markdown消息
func (ding *dingTalk) Markdown(title, context string, opt ...ding.AtOption) Markdown {
	initMarkDown := &markdown{
		title: title,
	}
	initMarkDown.name = ding.name
	initMarkDown.context = context
	initMarkDown.opt = opt
	return initMarkDown
}

//发送Text消息
func (ding *dingTalk) Text(context string, opt ...ding.AtOption) Text {
	initText := &text{}
	initText.name = ding.name
	initText.context = context
	initText.opt = opt
	return initText
}

//发送Text消息
func (ding *dingTalk) LinkMessage(title string, text string, picUrl string, msgUrl string) LinkMessage {
	initText := &uRLLink{}
	initText.name = ding.name
	initText.title = title
	initText.text = text
	initText.picUrl = picUrl
	initText.msgUrl = msgUrl
	return initText
}

//发送
func (ding *uRLLink) Send() error {
	if talk, ok := initDingTalk[ding.name]; ok {
		return talk.SendLinkMessage(ding.title, ding.text, ding.picUrl, ding.msgUrl)
	}
	return errors.New("类型错误")
}

//发送
func (ding *text) Send() error {
	if talk, ok := initDingTalk[ding.name]; ok {
		return talk.SendTextMessage(ding.context, ding.opt...)
	}
	return errors.New("类型错误")
}

//发送
func (ding *markdown) Send() error {
	if talk, ok := initDingTalk[ding.name]; ok {
		return talk.SendMarkDownMessage(ding.title, ding.context, ding.opt...)
	}
	return errors.New("类型错误")
}

//MarkDown格式处理
//换行符号，颜色，链接， 图片
//字体: 换行符，颜色，链接
//图片: 换行符? 链接?
