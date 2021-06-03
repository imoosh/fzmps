package api

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/conf"
    "centnet-fzmps/dao"
    "fmt"
    "github.com/gin-gonic/gin"
)

var srv *Service

type ginLogger struct {
}

func (l ginLogger) Write(p []byte) (n int, err error) {
    log.Debug(string(p))
    return 0, nil
}

type Service struct {
    router *gin.Engine
    dao    *dao.Dao
}

func Init(d *dao.Dao) {

    srv = &Service{
        dao:    d,
        router: routeInit(d),
    }

    var api = conf.Conf.API
    srv.router.Run(fmt.Sprintf("%s:%d", api.Host, api.Port))
}
