package ding

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gck/utils"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type DingTalk struct {
	robotToken string
	secret     string
}

func InitDingTalk(token string) *DingTalk {
	if token == "" {
		panic("no token")
	}
	return &DingTalk{
		robotToken: token,
	}
}

func InitDingTalkWithSecret(token string, secret string) *DingTalk {
	if token == "" || secret == "" {
		panic("no token")
	}
	return &DingTalk{
		robotToken: token,
		secret:     secret,
	}
}

func (d *DingTalk) sendMessage(msg iDingMsg) error {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		uri    string
		resp   *http.Response
		err    error
	)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	value := url.Values{}
	value.Set("access_token", d.robotToken)
	if d.secret != "" {
		t := time.Now().UnixNano() / 1e6
		value.Set("timestamp", fmt.Sprintf("%d", t))
		value.Set("sign", d.sign(t, d.secret))

	}
	uri = dingTalkURL + value.Encode()
	header := map[string]string{
		"Content-type": "application/json",
	}
	resp, err = utils.Request(ctx, "POST", uri, header, msg.Marshaler())

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("send msg err: %s, token: %s, msg: %s", string(body), d.robotToken, msg.Marshaler())
	}
	return nil
}

func (d *DingTalk) sign(t int64, secret string) string {
	strToHash := fmt.Sprintf("%d\n%s", t, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}

func (d *DingTalk) OutGoing(r io.Reader) (outGoingMsg OutGoingModel, err error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, &outGoingMsg)
	return
}

func (d *DingTalk) SendTextMessage(content string, opt ...AtOption) error {
	return d.sendMessage(NewTextMsg(content, opt...))
}

func (d *DingTalk) SendMarkDownMessage(title string, text Format, opts ...AtOption) error {
	return d.sendMessage(NewMarkDownMsg(title, text.GetContext(), opts...))
}

// 利用dtmd发送点击消息
func (d *DingTalk) SendDTMDMessage(title string, dtmdMap *dingMap, opt ...AtOption) error {
	return d.sendMessage(NewDTMDMsg(title, dtmdMap, opt...))
}

func (d DingTalk) SendMarkDownMessageBySlice(title string, textList []string, opts ...AtOption) error {
	text := ""
	for _, t := range textList {
		text = text + "\n" + t
	}
	return d.sendMessage(NewMarkDownMsg(title, text, opts...))
}

func (d *DingTalk) SendLinkMessage(title, text, picUrl, msgUrl string) error {
	return d.sendMessage(NewLinkMsg(title, text, picUrl, msgUrl))
}

func (d *DingTalk) SendActionCardMessage(title, text string, opts ...ActionCardOption) error {
	return d.sendMessage(NewActionCardMsg(title, text, opts...))
}

func (d *DingTalk) SendActionCardMessageBySlice(title string, textList []string, opts ...ActionCardOption) error {
	text := ""
	for _, t := range textList {
		text = text + "\n >" + t
	}
	return d.sendMessage(NewActionCardMsg(title, text, opts...))
}

func (d *DingTalk) SendFeedCardMessage(feedCard []FeedCardLinkModel) error {
	return d.sendMessage(NewFeedCardMsg(feedCard))
}
