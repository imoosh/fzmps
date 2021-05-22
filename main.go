package main

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/conf"
    "centnet-fzmps/dao"
    "centnet-fzmps/routers"
    "centnet-fzmps/services"
    "flag"
    "fmt"
    "runtime"

    _ "centnet-fzmps/models"
    _ "centnet-fzmps/routers"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    /* 解析参数 */
    flag.Parse()
    conf.Init()
    fmt.Println(conf.Conf)

    log.Init(conf.Conf.Logging)

    dao, err := dao.New(conf.Conf)
    if err != nil {
        panic(err)
    }

    services.Init(routers.Init(dao))
}
