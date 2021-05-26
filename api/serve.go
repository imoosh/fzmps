package api

import (
    "centnet-fzmps/conf"
    "centnet-fzmps/dao"
    "fmt"
    "github.com/gin-gonic/gin"
)

var srv *Service

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
