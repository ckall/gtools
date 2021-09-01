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

	//AddTextUrl(text string, hrefs []string)
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

func (f *format) GetContext() string {
	return strings.Join(f.text, separator)
}

//func (f *format) AddTextUrl(text string) {
//	var href string
//	for i := 0; i < len(hrefs); i += 2 {
//		href += fmt.Sprintf(urlFormat, hrefs[i], hrefs[i+1])
//	}
//	context := text + ":" + href
//	f.text = append(f.text, context)
//}

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
