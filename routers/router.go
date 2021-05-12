package routers

import (
	"centnet-fzmp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("weChat/v2/login", &controllers.LoginController{}, "post:Login")
}
