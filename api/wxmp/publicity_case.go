package wxmp

import (
    "centnet-fzmps/utils"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

type PublicityCaseResp struct {
    Id            string `json:"id"`
    Title         string `json:"title"`
    Picture       string `json:"picture"`
    Content       string `json:"content"`
    CaseType      string `json:"case_type"`
    Source        string `json:"source"`
    PublisherId   string `json:"publisher_id"`
    PublisherName string `json:"publisher_name"`
    PublisherTime string `json:"publisher_time"`
}

var CaseTypes = map[int]string{
    1:  "带宽、代办信用卡类",
    2:  "刷单返利类",
    4:  "虚假购物、服务类",
    5:  "杀猪盘类",
    6:  "冒充公检法及政府机关类",
    7:  "冒充领导、熟人等身份类",
    8:  "网络游戏产品虚假交易类",
    9:  "网络婚恋、交友类",
    10: "虚假征信类",
    11: "冒充军警购物类",
    12: "其他类型",
    13: "刷单返利类",
}

func RequestPublicityCase(c *gin.Context) {
    openid := c.Query("openid")
    if len(openid) == 0 {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
        return
    }

    var resp []PublicityCaseResp
    pc := DB(c).QueryPublicityCase()
    for _, c := range pc {
        resp = append(resp, PublicityCaseResp{
            Id:            fmt.Sprintf("%d", c.ID),
            Title:         c.Title,
            Picture:       c.Picture,
            Content:       c.Content,
            CaseType:      CaseTypes[c.CaseType],
            Source:        c.Source,
            PublisherId:   fmt.Sprintf("%d", c.PublisherId),
            PublisherName: c.PublisherName,
            PublisherTime: c.PublisherTime.Time().Format("2006-01-02 15:04:05"),
        })
    }

    c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&resp))
}

func ConfirmPublicityCase(c *gin.Context) {
    openid := c.Query("openid")
    if len(openid) == 0 {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
        return
    }
    caseId := c.Query("caseId")
    if len(caseId) == 0 {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
        return
    }

    c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
}
