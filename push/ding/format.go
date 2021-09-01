package ding

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	separator   = "\n >"
	imageFormat = "![screenshot](%s)"
	urlFormat   = "[%s](%s)"
)

type Format interface {
	//添加一行
	AddText(text string)
	//添加图片
	AddImage(url string)
	//一次性加多个图片
	AddImages(urls []string)
	//获取处理好的数据格式
	GetContext() string
	//以key：value的数据格式
	AddKeyValue(title string, value interface{})
	//链接
	AddTextUrl(text string, hrefs map[string]string)
	//H1字段，并新增颜色模块
	AddTextH1(context string, color ...interface{})
	//H2字段，并新增颜色模块
	AddTextH2(context string, color ...interface{})
	//H3字段，并新增颜色模块
	AddTextH3(context string, color ...interface{})
	//H4字段，并新增颜色模块
	AddTextH4(context string, color ...interface{})
	//H5字段，并新增颜色模块
	AddTextH5(context string, color ...interface{})
	//H6字段，并新增颜色模块
	AddTextH6(context string, color ...interface{})
}

//
func AddRed(text string) string {
	return fmt.Sprintf(hMap[RED], text)
}

//
func AddBlue(text string) string {
	return fmt.Sprintf(hMap[BLUE], text)
}

//
func AddGreen(text string) string {
	return fmt.Sprintf(hMap[GREEN], text)
}

//
func AddGold(text string) string {
	return fmt.Sprintf(hMap[GOLD], text)
}

/**
 * @auth: kuncheng
 * @Date: 2021/9/1
 */
type format struct {
	text []string
}

func NewConText() Format {
	return &format{}
}

func (f *format) AddText(text string) {
	f.text = append(f.text, text)
}

func (f *format) AddImage(url string) {
	f.text = append(f.text, fmt.Sprintf(imageFormat, url))
}

func (f *format) AddImages(urls []string) {
	for i := 0; i < len(urls); i++ {
		f.AddImage(urls[i])
	}
}

//输出原生语句可以自己研究
func (f *format) GetContext() string {
	return strings.Join(f.text, separator)
}

func (f *format) AddTextUrl(text string, hrefs map[string]string) {
	var href []interface{}
	for title, url := range hrefs {
		href = append(href, fmt.Sprintf(urlFormat, title, url))
	}
	var context = fmt.Sprintf(text, href...)
	f.text = append(f.text, context)
}

func (f *format) AddKeyValue(title string, value interface{}) {
	var str string
	switch value.(type) {
	case string:
		str = value.(string)
	default:
		b, _ := json.Marshal(value)
		str = string(b)
	}
	f.text = append(f.text, title+str)
}

func (f *format) AddTextH1(context string, color ...interface{}) {
	f.text = append(f.text, fmt.Sprintf(fmt.Sprintf(hMap[H1], context), color...))
}

func (f *format) AddTextH2(context string, color ...interface{}) {
	f.text = append(f.text, fmt.Sprintf(fmt.Sprintf(hMap[H2], context), color...))
}

func (f *format) AddTextH3(context string, color ...interface{}) {
	f.text = append(f.text, fmt.Sprintf(fmt.Sprintf(hMap[H3], context), color...))
}

func (f *format) AddTextH4(context string, color ...interface{}) {
	f.text = append(f.text, fmt.Sprintf(fmt.Sprintf(hMap[H4], context), color...))
}

func (f *format) AddTextH5(context string, color ...interface{}) {
	f.text = append(f.text, fmt.Sprintf(fmt.Sprintf(hMap[H5], context), color...))
}

func (f *format) AddTextH6(context string, color ...interface{}) {
	f.text = append(f.text, fmt.Sprintf(fmt.Sprintf(hMap[H6], context), color...))
}
