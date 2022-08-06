package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isther/easyTrader/conf"
	"github.com/isther/easyTrader/controller"
	"github.com/isther/easyTrader/middleware"
	"github.com/isther/easyTrader/pkg/jwt"
	"github.com/sirupsen/logrus"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Logger(logrus.StandardLogger()), gin.Recovery())

	var jwtAuth *gin.RouterGroup
	if conf.Conf.Dev {
		jwtAuth = router.Group("/api")
	} else {
		jwtAuth = router.Group("/api", middleware.JWTAuth())
	}

	router.GET("/api/ping", ping)

	//{{{ User routers
	var userGroup = router.Group("/api/user")
	{
		userGroup.POST("/register", controller.NewUserController().Create)
		userGroup.POST("/login", controller.NewUserController().Login)
	}

	var userAuthGroup = jwtAuth.Group("/user")
	{
		userAuthGroup.POST("/hello", hello)
		userAuthGroup.POST("/set/password", controller.NewUserController().SetPassword)
		userAuthGroup.POST("/set/binance/apiKey", controller.NewUserController().SetBinanceApiKey)
		userAuthGroup.POST("/set/binance/secretKey", controller.NewUserController().SetBinanceSecretKey)
		userAuthGroup.POST("/set/dingdingTalk/accessToken", controller.NewUserController().SetDingDingTalkAccessToken)
		userAuthGroup.POST("/set/dingdingTalk/secret", controller.NewUserController().SetDingDingTalkSecret)
	}
	//}}}

	//{{{ Trader routers
	var traderAuthGroup = jwtAuth.Group("/trader")
	{
		traderAuthGroup.POST("/trade", nil)
	}
	//}}}

	return router
}

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func hello(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	claims, err := jwt.ParseMapClaimsJwt(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Hello, " + claims["username"].(string),
	})
}
