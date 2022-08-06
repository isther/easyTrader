package middleware

import (
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(logger logrus.FieldLogger, notLogged ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		clientIP := ctx.ClientIP()
		clientUserAgent := ctx.Request.UserAgent()
		referer := ctx.Request.Referer()
		L := logger.WithFields(logrus.Fields{
			"clientIP": clientIP,
			"method":   ctx.Request.Method,
			"path":     path,
			"referer":  referer,
		})
		ctx.Set("L", L)

		start := time.Now()
		ctx.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := ctx.Writer.Status()
		dataLength := ctx.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := L.WithFields(logrus.Fields{
			"statusCode": statusCode,
			"latency":    latency,
			"dataLength": dataLength,
		})

		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := clientUserAgent
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
