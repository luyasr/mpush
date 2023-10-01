package ding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/luyasr/mpush/app/channel"
	"github.com/luyasr/mpush/app/message"
	"net/http"
	"strings"
	"time"
)

type dingMessageRequest struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	MarkDown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

type dingMessageResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newDingMessageRequest(message *message.Message) *dingMessageRequest {
	req := &dingMessageRequest{}
	if message.Text == "" {
		req.MsgType = "text"
		req.Text.Content = message.Content
	} else {
		req.MsgType = "markdown"
		req.MarkDown.Title = message.Title
		req.MarkDown.Text = message.Text
	}

	if message.To != "" {
		if message.To == "@all" {
			req.At.IsAtAll = true
		} else {
			req.At.AtUserIds = strings.Split(strings.TrimSpace(message.To), ",")
		}
	}

	return req
}

func SendDingMessage(message *message.Message, channel_ *channel.Channel) error {
	messageRequest := newDingMessageRequest(message)
	timestamp := time.Now().UnixMilli()
	sign := dingSign(channel_.Secret, timestamp)
	jsonData, err := json.Marshal(messageRequest)
	if err != nil {
		return err
	}
	response, err := http.Post(
		fmt.Sprintf("%s&timestamp=%d&sign=%s", channel_.Url, timestamp, sign),
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return err
	}

	var resp dingMessageResponse
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return errors.New(resp.Message)
	}

	return nil
}

func dingSign(secret string, timestamp int64) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	signBytes := h.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(signBytes)

	return signature
}
