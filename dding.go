package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	WebHookUrl = "https://oapi.dingtalk.com/robot/send?access_token="

	// MessageType
	textType     = "text"
	markdownType = "markdown"
)

type Robot struct {
	AccessToken string
	Secret      string
}

type RobotMsgOption func(Message)
type Message map[string]interface{}

// NewRobot 新建机器人
func NewRobot(accessToken, secret string) *Robot {
	return &Robot{
		AccessToken: accessToken,
		Secret:      secret,
	}
}

// BuildMsgAndSend 构建信息
func (r *Robot) BuildMsgAndSend(opt RobotMsgOption) error {
	msgMap := Message{}
	opt(msgMap)

	return r.send(msgMap)
}

// 发送信息
func (r *Robot) send(msg Message) error {
	// 构造参数
	url := r.buildQuery()
	reqData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// 发送请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		return err
	}

	// 获取钉钉响应
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("respData: %+v\n", string(respData))
	return nil
}

func (r *Robot) buildQuery() string {
	timestamp := time.Now().UnixMilli()
	signature := fmt.Sprintf("%d\n%s", timestamp, r.Secret)

	// hmac加密
	hash := hmac.New(sha256.New, []byte(r.Secret))
	hash.Write([]byte(signature))
	sign := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	// 拼接Query
	webhook := WebHookUrl + r.AccessToken
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", webhook, timestamp, sign)
	return url
}

// WithText 添加文本信息
func WithText(content string) RobotMsgOption {
	return func(m Message) {
		m["msgtype"] = textType
		m[textType] = map[string]string{
			"content": content,
		}
	}
}

func WithMarkDown(title, text string) RobotMsgOption {
	return func(m Message) {
		m["msgtype"] = markdownType
		m[markdownType] = map[string]string{
			"title": title,
			"text":  text,
		}
	}
}
