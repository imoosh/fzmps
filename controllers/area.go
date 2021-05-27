package controllers

import (
	"centnet-fzmps/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AreaInfo struct {
}

func GetAreaInfo(c *gin.Context) {
	t := c.Param("t")
	if t == "0" {
		if provinces := DB(c).QueryProvinces(); provinces != nil {
			var prov []string
			for _, v := range provinces {
				prov = append(prov, v.ProvinceName)
			}
			c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(prov))
		}
	}

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
}

func GetCooperationOrgInfo(c *gin.Context) {
	t := c.Param("t")
	if t == "0" {
		if provinces := DB(c).QueryProvinces(); provinces != nil {
			var prov []string
			for _, v := range provinces {
				prov = append(prov, v.ProvinceName)
			}
			c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(prov))
		}
	}

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
}
