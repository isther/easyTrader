package dao

import (
	"context"

	"github.com/isther/easyTrader/conf"
	"github.com/qiniu/qmgo"
	"github.com/sirupsen/logrus"
)

var db *qmgo.Database

var (
	UserCol *qmgo.Collection
)

func init() {
	var err error
	client, err := qmgo.Open(context.Background(), &qmgo.Config{
		Uri:      conf.Conf.DB.MongoUri,
		Database: conf.Conf.DB.DbName,
	})
	if err != nil {
		logrus.WithError(err).Fatalln("failed to connect to db!")
	}
	db = client.Database
	UserCol = db.Collection("users")
}
