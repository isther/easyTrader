package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isther/easyTrader/model"
	"github.com/isther/easyTrader/pkg/jwt"
	"github.com/isther/easyTrader/service"
	"github.com/sirupsen/logrus"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Create(ctx *gin.Context) {
	var (
		user model.User
		err  error
		L    = ctx.Value("L").(*logrus.Entry)
	)
	if err = ctx.ShouldBindJSON(&user); err != nil {
		L.WithError(err).Errorln("failed to bind json")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err = service.NewUserService().CreateUser(&user); err != nil {
		L.WithError(err).Errorln("failed to create new user")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (u *UserController) Login(ctx *gin.Context) {
	var (
		user model.User
		err  error
		L    = ctx.Value("L").(*logrus.Entry)
	)
	if err = ctx.ShouldBindJSON(&user); err != nil {
		L.WithError(err).Errorln("failed to bind json")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	token, err := service.NewUserService().Login(&user)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"token": token,
	})
}

func (u *UserController) SetPassword(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = make(map[string]string)
	)

	token := ctx.GetHeader("Authorization")
	claims, err := jwt.ParseMapClaimsJwt(token)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := model.User{
		Username: claims["username"].(string),
	}

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	password := tmpParams["password"]

	err = service.NewUserService().SetPassword(&user, password)
	if err != nil {
		L.WithError(err).Errorln("failed to set password")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (u *UserController) SetBinanceApiKey(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = make(map[string]string)
	)

	token := ctx.GetHeader("Authorization")
	claims, err := jwt.ParseMapClaimsJwt(token)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := model.User{
		Username: claims["username"].(string),
	}

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	apiKey := tmpParams["apiKey"]

	err = service.NewUserService().SetBinanceApiKey(&user, apiKey)
	if err != nil {
		L.WithError(err).Errorln("failed to set apiKey")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (u *UserController) SetBinanceSecretKey(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = make(map[string]string)
	)

	token := ctx.GetHeader("Authorization")
	claims, err := jwt.ParseMapClaimsJwt(token)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := model.User{
		Username: claims["username"].(string),
	}

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	secretKey := tmpParams["secretKey"]

	err = service.NewUserService().SetBinanceSecretKey(&user, secretKey)
	if err != nil {
		L.WithError(err).Errorln("failed to set secretKey")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (u *UserController) SetDingDingTalkAccessToken(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = make(map[string]string)
	)

	token := ctx.GetHeader("Authorization")
	claims, err := jwt.ParseMapClaimsJwt(token)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := model.User{
		Username: claims["username"].(string),
	}

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	accessToken := tmpParams["accessToken"]

	err = service.NewUserService().SetDingDingTalkAccessToken(&user, accessToken)
	if err != nil {
		L.WithError(err).Errorln("failed to set apiKey")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (u *UserController) SetDingDingTalkSecret(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = make(map[string]string)
	)

	token := ctx.GetHeader("Authorization")
	claims, err := jwt.ParseMapClaimsJwt(token)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := model.User{
		Username: claims["username"].(string),
	}

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	secret := tmpParams["secret"]

	err = service.NewUserService().SetDingDingTalkSecret(&user, secret)
	if err != nil {
		L.WithError(err).Errorln("failed to set secretKey")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (u *UserController) SetSymbols(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = make(map[string][]string)
	)

	token := ctx.GetHeader("Authorization")
	claims, err := jwt.ParseMapClaimsJwt(token)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := model.User{
		Username: claims["username"].(string),
	}

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// symbols := tmpParams["symbols"].([]string)
	symbols := tmpParams["symbols"]

	err = service.NewUserService().SetSymbols(&user, symbols...)
	if err != nil {
		L.WithError(err).Errorln("failed to set secretKey")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
