package api

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/controllers"
    "centnet-fzmps/dao"
    "github.com/gin-gonic/gin"
    "net/http"
)

type ginLogger struct {
}

func (l ginLogger) Write(p []byte) (n int, err error) {
    log.Debug(string(p))
    return 0, nil
}

func routeInit(dao *dao.Dao) *gin.Engine {
    router := gin.New()
    router.Use(gin.LoggerWithWriter(&ginLogger{}), gin.Recovery(), func(c *gin.Context) { c.Set("db", dao) })

    // 通过微信登陆
    router.POST("/api/wechat/login", controllers.WeChatLogin)

    // 更新数据
    router.POST("/api/wechat/update", controllers.UpdateRegisterInfo)

    // 完善注册信息
    router.POST("/api/wechat/perfect", controllers.PerfectRegisterInfo)

    // 获取用户信息
    router.GET("/api/user/get", controllers.RequestUserInfo)

    // 更新用户信息
    router.POST("/api/user/update", controllers.UpdateUserInfo)

    // 获取字典信息
    router.POST("/api/dict/list", controllers.RequestDictList)

    // 获取家人列表
    router.GET("/api/relative/list", controllers.RequestFamilyMembersList)

    // 删除家人信息
    router.POST("/api/relative/delete/:id", controllers.DeleteFamilyMember)

    // 添加家人信息
    router.POST("/api/relative/add", controllers.AddFamilyMember)

    // 更新家人信息
    router.PUT("/api/relative/udpate", controllers.UpdateFamilyMember)

    // 获取家人信息
    router.GET("/api/relative/:id", controllers.RequestFamilyMember)

    // 守护详情列表
    router.POST("/api/record/list", controllers.RequestRecordList)

    // 守护详情饼状图数据
    router.POST("/api/guard/count/pie", controllers.RequestRecordPie)

    router.GET("/api/wechat/sysArea/getAreaInfo/:t", controllers.GetAreaInfo)

    router.GET("/api/wechat/sysArea/org/list/:t", controllers.GetCooperationOrgInfo)

    router.GET("/api/alarm/get", controllers.RequestAlarm)

    router.GET("/api/alarm/confirm", controllers.ConfirmAlarm)

    // 针对企业微信后台开放接口
    wecom := router.Group("/api/wecom")
    {
        wecom.POST("/alarm", WeComVictimAlarm)
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
