package routers

import (
    "centnet-fzmps/common/log"
    "github.com/gin-gonic/gin"
)

type logger struct {
}

func (l logger) Write(p []byte) (n int, err error) {
    log.Debug(string(p))
    return 0, nil
}

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        gin.LoggerWithWriter(&logger{})
    }
}
