package ding

import "strings"

type Format interface {
	AddText(text string)
	GetContext() string
}

/**
 * @auth: kuncheng
 * @Date: 2021/9/1
 */
type format struct {
	text []string
}

func NewConText() *format {
	return &format{}
}

func (c *format) AddText(text string) {
	c.text = append(c.text, text)
}

func (c *format) GetContext() string {
	return strings.Join(c.text, "\n >")
}
