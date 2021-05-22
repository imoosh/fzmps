package controllers

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/utils"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

type DictListReq struct {
    Type string `json:"type"`
}

var (
    agePhases = []DictItem{
        {Label: "18-30岁", Value: "1"},
        {Label: "30-45岁", Value: "2"},
        {Label: "45-60岁", Value: "3"},
        {Label: "60岁以上", Value: "4"},
    }
    contacts = []DictItem{
        {Label: "夫妻", Value: "1"},
        {Label: "父母", Value: "2"},
        {Label: "子女", Value: "3"},
        {Label: "爷爷奶奶", Value: "4"},
        {Label: "兄弟姐妹", Value: "5"},
        {Label: "其他", Value: "6"},
    }
    professionalList = []DictItem{
        {Label: "其他职业", Value: "1"},
        {Label: "普通员工", Value: "2"},
        {Label: "个体户", Value: "3"},
        {Label: "军政人员", Value: "4"},
        {Label: "家庭主妇", Value: "5"},
        {Label: "退休人员", Value: "6"},
        {Label: "待业人员", Value: "7"},
        {Label: "学生党", Value: "8"},
    }
)

type DictItem struct {
    Label string `json:"label"`
    Value string `json:"value"`
}

// api/dict/list：获取字典信列表（年龄、职业、关联家人列表）
func GetDictList(c *gin.Context) {
    var req DictListReq

    err := c.ShouldBindJSON(&req)
    if err != nil {
        log.Error(err)
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
        return
    }

    switch strings.TrimSpace(req.Type) {
    case "age_group":
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&agePhases))
    case "profession":
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&professionalList))
    case "relation":
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&contacts))
    default:
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
    }
}
