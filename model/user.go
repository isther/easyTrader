package model

import (
	"context"

	"github.com/isther/easyTrader/dao"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`

	Binance struct {
		ApiKey    string `json:"apiKey" bson:"apiKey"`
		SecretKey string `json:"secretKey" bson:"secretKey"`
	} `json:"binance" bson:"binance"`

	DingDingTalk struct {
		ApiKey    string `json:"apiKey" bson:"apiKey"`
		SecretKey string `json:"secretKey" bson:"secretKey"`
	} `json:"dingdingTalk" bson:"dingdingTalk"`
	// ...
}

// user -> n * symbol -> n * config

func NewUser(pUser *User) error {
	rst, err := dao.UserCol.InsertOne(context.Background(), pUser)
	if err != nil {
		return err
	}
	logrus.Infof("user %v[%v] created!", pUser.Username, rst.InsertedID)
	return nil
}

func UsernameExist(username string) bool {
	return dao.UserCol.Find(context.Background(), bson.M{"username": username}).One(&User{}) == nil
}

func FindUserByUsername(username string) (*User, error) {
	uf := &User{}
	err := dao.UserCol.Find(context.Background(), bson.M{"username": username}).One(uf)
	if err != nil {
		return nil, err
	}
	return uf, nil
}

// func FindBinanceKeyByUsername(username string)

func (u *User) SetPassword(password string) error {
	err := dao.UserCol.UpdateOne(
		context.Background(),
		bson.M{"username": u.Username},
		bson.M{"$set": bson.M{"password": password}},
	)
	u.Password = password
	return err
}

func (u *User) SetBinanceApiKey(apiKey string) error {
	err := dao.UserCol.UpdateOne(
		context.Background(),
		bson.M{"username": u.Username},
		bson.M{"$set": bson.M{"binance.apiKey": apiKey}},
	)
	u.Password = apiKey
	return err
}

func (u *User) SetBinanceSecretKey(secretKey string) error {
	err := dao.UserCol.UpdateOne(
		context.Background(),
		bson.M{"username": u.Username},
		bson.M{"$set": bson.M{"binance.secretKey": secretKey}},
	)
	u.Password = secretKey
	return err
}

func (u *User) SetDingDingTalkApiKey(apiKey string) error {
	err := dao.UserCol.UpdateOne(
		context.Background(),
		bson.M{"username": u.Username},
		bson.M{"$set": bson.M{"dingdingTalk.apiKey": apiKey}},
	)
	u.Password = apiKey
	return err
}

func (u *User) SetDingDingTalkSecretKey(secretKey string) error {
	err := dao.UserCol.UpdateOne(
		context.Background(),
		bson.M{"username": u.Username},
		bson.M{"$set": bson.M{"dingdingTalk.secretKey": secretKey}},
	)
	u.Password = secretKey
	return err
}
