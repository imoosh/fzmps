package async

import (
    "context"
    "errors"
    "github.com/hibiken/asynq"
)

type AsyncProc struct {
    cli *asynq.Client
    srv *asynq.Server
    mux *asynq.ServeMux
}

type Config struct {
    Addr string
    Auth string
}

type HandleFunc func(context.Context, *asynq.Task) error

func NewAsyncProc(c Config) (acp *AsyncProc, err error) {

    opt := asynq.RedisClientOpt{Addr: c.Addr, Password: c.Auth}
    cfg := asynq.Config{Logger: logger{}, LogLevel: defaultLoggerLevel}
    srv := asynq.NewServer(opt, cfg)
    cli := asynq.NewClient(opt)

    acp = &AsyncProc{
        cli: cli,
        srv: srv,
        mux: asynq.NewServeMux(),
    }

    if acp.srv == nil || acp.cli == nil {
        return nil, errors.New("create async processor failed")
    }
    //acp.srv.Run(acp.mux)
    return
}

func (ap *AsyncProc) Run() {
    ap.srv.Run(ap.mux)
}

func (ap *AsyncProc) Register(pattern string, handler HandleFunc) {
    ap.mux.HandleFunc(pattern, handler)
}

/*
Enqueue(task, asynq.ProcessAt(t))
Enqueue(task, asynq.ProcessIn(t))
*/
func (ap *AsyncProc) Enqueue(task *asynq.Task, opts ...asynq.Option) (err error) {
    _, err = ap.cli.Enqueue(task, opts...)
    return
}
