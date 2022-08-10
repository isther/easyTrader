package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	KeepLiveTimeout = 60 * time.Second * 10
)

type WsController struct{}

func NewWsController() *WsController {
	return &WsController{}
}

func (ws *WsController) Ping(ctx *gin.Context) {
	var (
		err error
	)

	logrus.Info("GOGOGOWS")
	wsConn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Error(err)
		return
	}

	defer func() {
		wsConn.Close()
	}()

	for {
		wsConn.SetWriteDeadline(time.Now().Add(KeepLiveTimeout))

		msgType, msg, err := wsConn.ReadMessage()
		if err != nil {
			logrus.Error(err)
			break
		}

		if string(msg) == "ping" {
			msg = []byte("pong")
		}

		err = wsConn.WriteMessage(msgType, msg)
		if err != nil {
			logrus.Error(err)
			break
		}
	}
}
