package push

import (
	"errors"
	"gck/push/ding"
)

var (
	dingTalk map[string]*ding.DingTalk
)

/**
 * @auth: kuncheng
 * @Date: 2021/8/30
 */
//关键词初始化
//tokens 可以发送给多个群
//key 每个群的多个关键字验证
func InitKeyWordDing(name, tokens string, title string) {
	dingTalk[name] = ding.InitDingTalk(tokens, title)
}

//签名规则
func InitSignDing(name, token string, secret string) {
	dingTalk[name] = ding.InitDingTalkWithSecret(token, secret)
}

//发送Markdown消息
func SendMarkdown(name string, title, context string, opt ...ding.AtOption) error {
	if talk, ok := dingTalk[name]; ok {
		return talk.SendMarkDownMessage(title, context, opt...)
	}
	return errors.New("类型错误")
}

//发送Text消息
func SendText(name string, context string, opt ...ding.AtOption) error {
	if talk, ok := dingTalk[name]; ok {
		return talk.SendTextMessage(context, opt...)
	}
	return errors.New("类型错误")
}
