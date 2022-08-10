package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isther/easyTrader/model"
	"github.com/isther/easyTrader/pkg/jwt"
	"github.com/isther/easyTrader/service"
	"github.com/sirupsen/logrus"
)

type SymbolsController struct{}

func NewSymbolsController() *SymbolsController {
	return &SymbolsController{}
}

func (s *SymbolsController) CheckIsOK(ctx *gin.Context) {
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

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	symbol := tmpParams["symbol"]

	ok, err := service.NewSymbolsService().CheckIsOk(claims["username"].(string), symbol)
	if err != nil {
		L.WithError(err).Errorln("failed to set check symbol")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"res": ok,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"res": ok,
	})
}

func (s *SymbolsController) SetSymbols(ctx *gin.Context) {
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

	err = service.NewSymbolsService().SetSymbols(&user, symbols...)
	if err != nil {
		L.WithError(err).Errorln("failed to set symbols")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func (s *SymbolsController) GetSymbols(ctx *gin.Context) {
	var (
		err error
		L   = ctx.Value("L").(*logrus.Entry)
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

	symbols, err := service.NewSymbolsService().GetSymbols(claims["username"].(string))
	if err != nil {
		L.WithError(err).Errorln("failed to get symbols")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":     "ok",
		"symbols": symbols,
	})
}

func (s *SymbolsController) SearchSymbols(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = struct {
			Amplitue           json.Number `json:"amplitue"`
			AmplituePercentage json.Number `json:"amplituePercentage"`
		}{}
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

	err = ctx.ShouldBindJSON(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	amplitue, _ := tmpParams.Amplitue.Float64()
	amplituePercentage, _ := tmpParams.AmplituePercentage.Float64()

	res, err := service.NewSymbolsService().SearchSymbols(claims["username"].(string), amplitue, amplituePercentage)
	if err != nil {
		L.WithError(err).Errorln("failed to search symbols by amplitue")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"res": res,
	})
}
