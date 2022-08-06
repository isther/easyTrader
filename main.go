package main

import (
	"github.com/isther/easyTrader/conf"
	"github.com/isther/easyTrader/routers"
	"github.com/sirupsen/logrus"
)

func main() {
	r := routers.Init()
	logrus.Info("Server listen: ", conf.Conf.Listen)
	r.Run(conf.Conf.Listen)
}
