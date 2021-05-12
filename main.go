package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.Listen.HTTPAddr = ""
	beego.BConfig.Listen.HTTPPort = 8089

	beego.InsertFilter("/api/*", beego.BeforeExec, nil, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8080
}
