package ding

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func getDingHook(t *testing.T) *Ding {
	type Conf struct {
		AccessToken string `yaml:"access_token"`
		Secret      string `yaml:"secret"`
	}

	var conf  = new(Conf)

	yamlFileBytes, err := ioutil.ReadFile("ding_conf.yml")
	if err != nil {
		t.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFileBytes, &conf)
	if err != nil {
		t.Fatal(err)
	}

	return NewDing(conf.AccessToken,conf.Secret)
}

func TestDingNormalMsg(t *testing.T) {
	getDingHook(t).SendMessage("普通消息测试")
}

func TestDingAtSomeMember(t *testing.T) {
	getDingHook(t).EnableAt().SendMessage("@成员消息测试", "")
}

func TestDingAtAllMember(t *testing.T) {
	getDingHook(t).EnableAt().EnbaleAtAll().SendMessage("@全体成员消息测试")
}
