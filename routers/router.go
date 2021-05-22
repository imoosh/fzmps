package routers

import (
    "centnet-fzmps/controllers"
    "centnet-fzmps/dao"
    "centnet-fzmps/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Init(dao *dao.Dao) *gin.Engine {
    router := gin.New()
    router.Use(gin.LoggerWithWriter(&GinLogger{}), gin.Recovery(), func(c *gin.Context) {
        c.Set("db", dao)
    })

    // 通过微信登陆
    router.POST("/api/wechat/login", controllers.WeChatLogin)

    // 更新数据
    router.POST("/api/wechat/update", controllers.UpdateRegisterInfo)

    // 完善注册信息
    router.POST("/api/wechat/perfect", controllers.PerfectRegisterInfo)

    // 获取用户信息
    router.GET("/api/user/get", controllers.GetUserInfo)

    // 更新用户信息
    router.POST("/api/user/update", controllers.UpdateUserInfo)

    // 获取家人相关信息
    router.POST("/api/dict/list", controllers.GetDictList)

    // 获取家人列表
    router.POST("/api/relative/getRelativeList", controllers.GetRelativeList)

    // 删除家人信息
    router.POST("api/relative/deleteRelative", controllers.DeleteRelative)

    // 添加家人信息
    router.POST("/api/relative/putRelative", controllers.AddRelative)

    // 更新家人信息
    router.POST("/api/relative/updateRelative", controllers.UpdateRelative)

    // 获取家人信息
    router.POST("/api/relative/get", controllers.GetRelative)

    // 守护详情列表
    router.POST("/api/record/list", controllers.GetRecordList)

    // 守护详情饼状图数据
    router.POST("/api/guard/count/pie", controllers.GetRecordPie)

    router.GET("/api/wechat/sysArea/getAreaInfo/:t", controllers.GetAreaInfo)

    router.GET("/api/wechat/sysArea/org/list/:t", controllers.GetCooperationOrgInfo)

    wecom := router.Group("/api/wecom")
    {
        wecom.POST("/api/wecom/pushed-alarm", services.WeComVictimAlarm)
    }

    authorized := router.Group("/")
    authorized.Use(AuthRequired())
    {

    }

    return router
}

//https://blog.csdn.net/qq_37767455/article/details/104712028
// https://juejin.cn/post/6844904090690912264
func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.Request.Header.Get("token")
        return
        if len(token) == 0 {
            c.JSON(http.StatusOK, gin.H{
                "ErrCode": -1,
                "ErrMsg":  "请求未携带token，无权限访问",
                "Data":    nil,
            })
            c.Abort()
            return
        } else {
            c.Next()
        }

    }
}
