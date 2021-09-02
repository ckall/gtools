package push

import (
	"errors"
	"github.com/ckall/gtools/push/ding"
)

type DingTalk interface {
	//Markdown模版类型
	Markdown(title string, context string, opt ...ding.AtOption) Markdown
	//Text模版类型
	Text(context string, opt ...ding.AtOption) Text
}

//链接模版类型接口访问
type LinkMessage interface {
	Send() error
}

//Markdown模版类型接口访问
type Markdown interface {
	Send() error
}

//文本类型接口访问
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
	title   string
	context string
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
func (ding *dingTalk) Markdown(title string, context string, opt ...ding.AtOption) Markdown {
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
func (ding *dingTalk) Link(title string, text string, picUrl string, msgUrl string) LinkMessage {
	initText := &uRLLink{}
	initText.name = ding.name
	initText.title = title
	initText.text = text
	initText.picUrl = picUrl
	initText.msgUrl = msgUrl
	return initText
}

//发送链接
func (ding *uRLLink) Send() error {
	if talk, ok := initDingTalk[ding.name]; ok {
		return talk.SendLinkMessage(ding.title, ding.text, ding.picUrl, ding.msgUrl)
	}
	return errors.New("类型错误")
}

//发送文本内容
func (ding *text) Send() error {
	if talk, ok := initDingTalk[ding.name]; ok {
		return talk.SendTextMessage(ding.context, ding.opt...)
	}
	return errors.New("类型错误")
}

//发送markdown内容
func (ding *markdown) Send() error {
	if talk, ok := initDingTalk[ding.name]; ok {
		return talk.SendMarkDownMessage(ding.title, ding.context, ding.opt...)
	}
	return errors.New("类型错误")
}
