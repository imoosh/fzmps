package routers

import (
	"centnet-fzmps/controllers"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	beego.Router("api/wechat/login", &controllers.LoginController{}, "post:Login")
	beego.Router("api/wechat/update", &controllers.LoginController{}, "post:Update")

	beego.Router("api/user", &controllers.UserController{}, "get:UserInfo")
	beego.Router("api/user/update", &controllers.UserController{}, "post:UpdateUserInfo")

	beego.Router("api/dict/list", &controllers.UserController{}, "post:DictList")
}

func init1() {
	router := gin.Default()
	router.POST("/api/wechat/login", controllers.WeChatLogin)
	router.POST("/api/wechat/update", controllers.UpdateRegisterInfo)

	router.GET("/api/user", controllers.GetUserInfo)
	router.POST("/api/user/update", controllers.UpdateUserInfo)
	router.POST("/api/dict/list", controllers.GetDictList)

	authorized := router.Group("/")
	authorized.Use(AuthRequired())
	{

	}
}

// https://juejin.cn/post/6844904090690912264
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"ErrCode": -1,
				"ErrMsg":  "请求未携带token，无权限访问",
				"Data":    nil,
			})
			c.Abort()
			return
		}

	}
}
