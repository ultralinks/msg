package org

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"msg/userLogic/service/model"
	orgService "msg/userLogic/service/org"
	orgUserService "msg/userLogic/service/orgUser"
	"msg/userLogic/util"
)

type CreateOrgRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type OrgResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

//
// @Summary 创建组织
// @Tags    组织
// @Accept  json
// @Produce  json
// @Param body body org.CreateOrgRequest true "请求参数"
// @Success 200 {string} string "{"msg": "ok"}"
// @Failure 400 {string} json "{"error": "error info"}"
// @Failure 500 {string} json "{"error": "error info"}"
// @Router /org [post]
func Create(c *gin.Context) {
	r := &CreateOrgRequest{}
	if err := util.ParseRequest(c, r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	org := model.Org{
		Id:      util.GetRandomString(11),
		Name:    r.Name,
		Desc:    r.Desc,
		Created: time.Now(),
		Updated: time.Now(),
	}
	if err := orgService.Create(&org); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetHeader("User_id")
	orgUser := model.OrgUser{
		OrgId:  org.Id,
		UserId: userId,
	}
	if err := orgUserService.Create(&orgUser); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orgId": org.Id})
}

//
// @Summary 获取组织详情
// @Tags    组织
// @Accept  json
// @Produce  json
// @Param id path string true "orgId"
// @Success 200 {string} json "{"msg": "ok"}"
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /org/{id} [get]
func Get(c *gin.Context) {
	orgId := c.Param("id")

	org, err := orgService.Get(orgId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02 15:04:05"
	orgResponse := OrgResponse{
		Id:      org.Id,
		Name:    org.Name,
		Desc:    org.Desc,
		Created: org.Created.Format(layout),
		Updated: org.Updated.Format(layout),
	}

	c.JSON(http.StatusOK, orgResponse)
}

//
// @Summary 获取组织列表
// @Tags    组织
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"msg": "ok"}"
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /orgs [get]
func List(c *gin.Context) {
	userId := c.GetHeader("User_id")
	fmt.Println("ListMethod, userId = ", userId)
	orgs, err := orgService.ListByUserId(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02 15:04:05"
	orgListResponse := make([]OrgResponse, 0)
	for _, o := range *orgs {
		orgListResponse = append(orgListResponse, OrgResponse{
			Id:      o.Id,
			Name:    o.Name,
			Desc:    o.Desc,
			Created: o.Created.Format(layout),
			Updated: o.Updated.Format(layout),
		})
	}

	c.JSON(http.StatusOK, orgListResponse)
}
