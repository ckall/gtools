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
	AddText(text string)

	AddImage(url string)

	AddImages(urls []string)

	GetContext() string

	AddKeyValue(title string, value interface{})

	AddTextUrl(text string, hrefs map[string]string)

	AddTextH1(context string, color ...interface{})

	AddTextH2(context string, color ...interface{})

	AddTextH3(context string, color ...interface{})

	AddTextH4(context string, color ...interface{})

	AddTextH5(context string, color ...interface{})

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
