package routers

import (
    "centnet-fzmps/conf"
    "centnet-fzmps/routers/api/wecom"
    "centnet-fzmps/routers/api/wxmp"
    "centnet-fzmps/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

var (
    srv *service.Service
)

func Init(c *conf.Config, s *service.Service) *gin.Engine {
    srv = s

    router := gin.New()
    router.Use(Logger())
    router.Use(gin.Recovery())

    // 通过微信登陆
    router.POST("/api/wechat/login", wxmp.WeChatLogin)

    // 更新数据
    router.POST("/api/wechat/update", wxmp.UpdateRegisterInfo)

    // 完善注册信息
    router.POST("/api/wechat/perfect", wxmp.PerfectRegisterInfo)

    // 获取用户信息
    router.GET("/api/user/get", wxmp.RequestUserInfo)

    // 更新用户信息
    router.POST("/api/user/update", wxmp.UpdateUserInfo)

    // 获取字典信息
    router.POST("/api/dict/list", wxmp.RequestDictList)

    // 获取家人列表
    router.GET("/api/relative/list", wxmp.RequestFamilyMembersList)

    // 删除家人信息
    router.POST("/api/relative/delete/:id", wxmp.DeleteFamilyMember)

    // 添加家人信息
    router.POST("/api/relative/add", wxmp.AddFamilyMember)

    // 更新家人信息
    router.PUT("/api/relative/update", wxmp.UpdateFamilyMember)

    // 获取家人信息
    router.GET("/api/relative/:id", wxmp.RequestFamilyMember)

    // 守护详情列表
    router.POST("/api/record/list", wxmp.RequestRecordList)

    // 守护详情饼状图数据
    router.POST("/api/guard/count/pie", wxmp.RequestRecordPie)

    //router.GET("/api/wechat/sysArea/getAreaInfo/:t", controllers.GetAreaInfo)
    //router.GET("/api/wechat/sysArea/org/list/:t", controllers.GetCooperationOrgInfo)

    // 请求预警消息
    router.GET("/api/alarm/get", wxmp.RequestAlarm)
    // 确认预警消息
    router.GET("/api/alarm/confirm", wxmp.ConfirmAlarm)

    // 请求宣传案例
    router.GET("/api/publicityCase/get", wxmp.RequestPublicityCase)
    // 获取案例正文
    router.GET("/api/publicityCaseContent/get", wxmp.RequestPublicityCaseContent)
    // 确认宣传案例
    router.GET("/api/publicityCase/confirm", wxmp.ConfirmPublicityCase)

    // 针对企业微信后台开放接口
    router.POST("/api/wecom/alarm", wecom.VictimAlarm)

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
