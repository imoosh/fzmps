package main

import (
    "centnet-fzmps/api"
    "centnet-fzmps/common/log"
    "centnet-fzmps/conf"
    "centnet-fzmps/dao"
    "flag"
    "fmt"
    "os"
    "os/signal"
    "runtime"
    "syscall"
    "time"
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

    api.Init(dao)

    //printMemStatsLoop()
    //debug.SetGCPercent(10)

    // os signal
    sigterm := make(chan os.Signal, 1)
    signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
    <-sigterm
}

func printMemStatsLoop(duration time.Duration) {
    var m runtime.MemStats
    ticker := time.NewTicker(duration)
    go func() {
        select {
        case <-ticker.C:
            runtime.ReadMemStats(&m)
            log.Debugf("Alloc = %vMB Sys = %vMB NumGC = %v",
                m.Alloc/1024/1024, m.Sys/1024/1024, m.NumGC)
        }
    }()
}
