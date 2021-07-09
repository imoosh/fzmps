package api

import (
    async "centnet-fzmps/common/aproc"
    "centnet-fzmps/common/log"
    "centnet-fzmps/conf"
    "centnet-fzmps/dao"
    "centnet-fzmps/models"
    "centnet-fzmps/routers"
    "context"
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/hibiken/asynq"
)

const (
    HandleWeComVictimAlarms = "handle-wecom-victim-alarms"
)

var srv *Service

type Service struct {
    router *gin.Engine
    dao    *dao.Dao
    proc   *async.AsyncProc
}

func New(d *dao.Dao, c *conf.Config) (*Service, error) {
    proc := async.NewAsyncProc(async.Config{
        Addr: c.Redis.Addr,
        Auth: c.Redis.Auth,
    })

    srv = &Service{
        dao:    d,
        proc:   proc,
        router: routers.Init(srv, d, proc),
    }

    var api = conf.Conf.API
    go srv.router.Run(fmt.Sprintf("%s:%d", api.Host, api.Port))

    return srv, nil
}

func (srv *Service) Run() {
    srv.proc.Register(HandleWeComVictimAlarms, srv.handleWeComVictimAlarms)

    srv.proc.Run()
}

func (srv *Service) handleWeComVictimAlarms(ctx context.Context, task *asynq.Task) error {
    var alarm models.FzmpsAlarm
    if err := json.Unmarshal(task.Payload(), &alarm); err != nil {
        log.Error(err)
        return nil
    }
    srv.dao.InsertWeComAlarm(&alarm)
    return nil
}

func (srv *Service) AsyncHandleWeComVictimAlarms(data *models.FzmpsAlarm) {
    bs, err := json.Marshal(data)
    if err != nil {
        log.Error(err)
        return
    }

    task := asynq.NewTask(HandleWeComVictimAlarms, bs)
    if err := srv.proc.Enqueue(task, asynq.MaxRetry(5)); err != nil {
        log.Error(err)
    }
}
