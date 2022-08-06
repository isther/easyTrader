package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isther/easyTrader/pkg/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := jwt.ParseMapClaimsJwtHeader(ctx)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusOK, err.Error())
			return
		}
		ctx.Next()
	}
}
