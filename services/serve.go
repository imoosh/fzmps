package services

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

func Init(r *gin.Engine) {

    srv = &Service{
        router: r,
    }

    var api = conf.Conf.MiniPro.API
    r.Run(fmt.Sprintf("%s:%d", api.Host, api.Port))

}
