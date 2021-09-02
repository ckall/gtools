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
}

//红色字体
func AddRed(text string) string {
	return fmt.Sprintf(hMap[red], text)
}

//蓝色
func AddBlue(text string) string {
	return fmt.Sprintf(hMap[blue], text)
}

//字体型号
func AddH1(text string) string {
	return fmt.Sprintf(hMap[h1], text)
}

//字体型号
func AddH2(text string) string {
	return fmt.Sprintf(hMap[h2], text)
}

//字体型号
func AddH3(text string) string {
	return fmt.Sprintf(hMap[h3], text)
}

//字体型号
func AddH4(text string) string {
	return fmt.Sprintf(hMap[h4], text)
}

//字体型号
func AddH5(text string) string {
	return fmt.Sprintf(hMap[h5], text)
}

//字体型号
func AddH6(text string) string {
	return fmt.Sprintf(hMap[h6], text)
}

//绿色
func AddGreen(text string) string {
	return fmt.Sprintf(hMap[green], text)
}

//黄金
func AddGold(text string) string {
	return fmt.Sprintf(hMap[gold], text)
}

//添加URL地址
func AddUrl(title string, url string) string {
	return fmt.Sprintf(urlFormat, title, url)
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

//添加文字
func (f *format) AddText(text string, color ...interface{}) {
	f.text = append(f.text, fmt.Sprintf(text, color...))
}

//添加图片
func (f *format) AddImage(url string) {
	f.text = append(f.text, fmt.Sprintf(imageFormat, url))
}

//一次性添加多个图片
func (f *format) AddImages(urls []string) {
	for i := 0; i < len(urls); i++ {
		f.AddImage(urls[i])
	}
}

//输出原生语句可以自己研究
func (f *format) GetContext() string {
	return strings.Join(f.text, separator)
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
