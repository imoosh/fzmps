package routers

import "centnet-fzmps/common/log"

type GinLogger struct {
}

func (l GinLogger) Write(p []byte) (n int, err error) {
    log.Debug(string(p))
    return 0, nil
}
