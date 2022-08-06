package ding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Ding struct {
	AccessToken string
	Secret      string

	At    bool
	AtAll bool
}

func NewDing(accessToken, secret string) *Ding {
	return &Ding{
		AccessToken: accessToken,
		Secret:      secret,
		At:    false,
		AtAll: false,
	}
}

func (d *Ding) EnableAt() *Ding {
	d.At = true
	return d
}

func (d *Ding) EnbaleAtAll() *Ding {
	d.AtAll = true
	return d
}

func (d *Ding) SendMessage(s string, at ...string) error {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": s,
		},
	}
	if d.At {
		if d.AtAll {
			if len(at) > 0 {
				return errors.New("the parameter \"AtAll\" is \"true\", but the \"at\" parameter of SendMessage is not empty")
			}
			msg["at"] = map[string]interface{}{
				"isAtAll": d.AtAll,
			}
		} else {
			msg["at"] = map[string]interface{}{
				"atMobiles": at,
				"isAtAll":   d.AtAll,
			}
		}
	} else {
		if len(at) > 0 {
			return errors.New("the parameter \"EnableAt\" is \"false\", but the \"at\" parameter of SendMessage is not empty")
		}
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(d.getURL(), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (d *Ding) hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (d *Ding) getURL() string {
	wh := "https://oapi.dingtalk.com/robot/send?access_token=" + d.AccessToken
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, d.Secret)
	sign := d.hmacSha256(stringToSign, d.Secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", wh, timestamp, sign)
	return url
}
