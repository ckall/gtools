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
	AddText(text string, color ...interface{})
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
}

//
func AddRed(text string) string {
	return fmt.Sprintf(hMap[red], text)
}

//
func AddBlue(text string) string {
	return fmt.Sprintf(hMap[blue], text)
}

//
func AddH1(text string) string {
	return fmt.Sprintf(hMap[h1], text)
}

//
func AddH2(text string) string {
	return fmt.Sprintf(hMap[h2], text)
}

//
func AddH3(text string) string {
	return fmt.Sprintf(hMap[h3], text)
}

//
func AddH4(text string) string {
	return fmt.Sprintf(hMap[h4], text)
}

//
func AddH5(text string) string {
	return fmt.Sprintf(hMap[h5], text)
}

//
func AddH6(text string) string {
	return fmt.Sprintf(hMap[h6], text)
}

//
func AddGreen(text string) string {
	return fmt.Sprintf(hMap[green], text)
}

//
func AddGold(text string) string {
	return fmt.Sprintf(hMap[gold], text)
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

func (f *format) AddText(text string, color ...interface{}) {
	f.text = append(f.text, fmt.Sprintf(text, color...))
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
	f.text = append(f.text, f.addTextUrl(text, hrefs))
}

func (f *format) addTextUrl(text string, hrefs map[string]string) string {
	var href []interface{}
	for title, url := range hrefs {
		href = append(href, fmt.Sprintf(urlFormat, title, url))
	}
	return fmt.Sprintf(text, href...)
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
