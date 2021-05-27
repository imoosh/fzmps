package controllers

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
			CaseType:      fmt.Sprintf("%d", c.CaseType),
			Source:        c.Source,
			PublisherId:   fmt.Sprintf("%d", c.PublisherId),
			PublisherName: c.PublisherName,
			PublisherTime: c.PublisherTime.Time().Format("2016-01-02 15:04:05"),
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
