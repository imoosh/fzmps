package wxmp

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
	"centnet-fzmps/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type familyMembersResp struct {
	Id           string `json:"id"`
	RelativeName string `json:"relativeName"`
	Relation     string `json:"relation"`
	Phone        string `json:"phone"`
}

type familyMember struct {
	Id           string `json:"id"`
	RelativeName string `json:"relativeName"`
	Relation     string `json:"relation"`
	Phone        string `json:"phone"`
	Sex          string `json:"sex"`
	Age          string `json:"age"`
	QQ           string `json:"qq"`
	WX           string `json:"wx"`
}

func RequestFamilyMembersList(c *gin.Context) {
	// 获取http头部token
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	// 通过token查询用户信息
	u, err := DB(c).QueryUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	var resp []familyMembersResp
	members := DB(c).QueryFamilyMembers(u.Id)
	for _, m := range members {
		resp = append(resp, familyMembersResp{
			Id:           fmt.Sprintf("%d", m.ID),
			RelativeName: m.Name,
			Relation:     m.Relation,
			Phone:        m.Phone,
		})
	}
	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&resp))
}

func DeleteFamilyMember(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	DB(c).DeleteFamilyMember(id)
	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
}

func AddFamilyMember(c *gin.Context) {
	var req familyMember
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
	}

	// 获取http头部token
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	// 通过token查询用户信息
	u, err := DB(c).QueryUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	m := &models.FzmpsFamily{
		UserId:    u.Id,
		Relation:  req.Relation,
		Name:      req.RelativeName,
		Phone:     req.Phone,
		Sex:       req.Sex,
		Age:       req.Age,
		QQ:        req.QQ,
		WX:        req.WX,
		IsDeleted: false,
	}
	DB(c).CreateFamilyMember(m)

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
}

func UpdateFamilyMember(c *gin.Context) {
	var req familyMember
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
	}

	// 获取http头部token
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	// 通过token查询用户信息
	u, err := DB(c).QueryUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	id, err := strconv.Atoi(req.Id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}
	m := &models.FzmpsFamily{
		ID:        uint(id),
		UserId:    u.Id,
		Relation:  req.Relation,
		Name:      req.RelativeName,
		Phone:     req.Phone,
		Sex:       req.Sex,
		Age:       req.Age,
		QQ:        req.QQ,
		WX:        req.WX,
		IsDeleted: false,
	}
	DB(c).UpdateFamilyMember(m)

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
}

func RequestFamilyMember(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	m := DB(c).QueryFamilyMember(id)
	if m == nil {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	tmp := &familyMember{
		Id:           id,
		RelativeName: m.Name,
		Relation:     m.Relation,
		Phone:        m.Phone,
		Sex:          m.Sex,
		Age:          m.Age,
		QQ:           m.QQ,
		WX:           m.WX,
	}

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(tmp))
}

type guardianRecord struct {
	TotalCount    int    `json:"totalCount"`
	UserId        string `json:"userId"`
	Name          string `json:"name"`
	Relation      string `json:"relation"`
	Phone         string `json:"phone"`
	FraudTime     string `json:"fraudTime"`
	FraudDuration string `json:"fraudDuration"`
}

func RequestRecordList(c *gin.Context) {
	// 获取http头部token
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	// 通过token查询用户信息
	u, err := DB(c).QueryUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	var rec []guardianRecord
	members := DB(c).QueryFamilyMembers(u.Id)
	for _, m := range members {
		rec = append(rec, guardianRecord{
			TotalCount:    0,
			UserId:        fmt.Sprintf("%d", u.Id),
			Name:          m.Name,
			Relation:      m.Relation,
			Phone:         m.Phone,
			FraudTime:     "",
			FraudDuration: "",
		})
	}
	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&rec))
}

type guardianRecordPie struct {
	Label   string `json:"label"`
	Numbers int    `json:"numbers"`
}

func RequestRecordPie(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&guardianRecordPie{}))
}
