package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"

	_ "centnet-fzmps/models"
	_ "centnet-fzmps/routers"
)

func main() {
	logger := logs.NewLogger()
	logger.SetLogger(logs.AdapterConsole)
	logs.EnableFuncCallDepth(true)

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.Listen.HTTPAddr = "192.168.1.14"
	beego.BConfig.Listen.HTTPPort = 6060

	beego.InsertFilter("/api/*", beego.BeforeExec, func(context *context.Context) {}, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8080
}
