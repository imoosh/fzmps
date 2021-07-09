package wxmp

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	agePhases  []DictItem
	contacts   []DictItem
	profession []DictItem
)

type DictListReq struct {
	Type string `json:"type"`
}

type DictItem struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// api/dict/list：获取字典信列表（年龄、职业、关联家人列表）
func RequestDictList(c *gin.Context) {
	var req DictListReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	switch strings.TrimSpace(req.Type) {
	case "age_group":
		if len(agePhases) == 0 {
			dict := DB(c).QueryAgePhasesDict()
			for _, i := range dict {
				agePhases = append(agePhases, DictItem{Label: i.Label, Value: i.Value})
			}
		}
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&agePhases))
	case "profession":
		if len(profession) == 0 {
			dict := DB(c).QueryProfessionDict()
			for _, i := range dict {
				profession = append(profession, DictItem{Label: i.Label, Value: i.Value})
			}
		}
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&profession))
	case "relation":
		if len(contacts) == 0 {
			dict := DB(c).QueryRelationDict()
			for _, i := range dict {
				contacts = append(contacts, DictItem{Label: i.Label, Value: i.Value})
			}
		}
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&contacts))
	default:
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
	}
}
