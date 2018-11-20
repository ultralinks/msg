package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	appService "msg/userLogic/service/app"
	"msg/userLogic/service/model"
	"msg/userLogic/util"
)

type CreateAppRequest struct {
	OrgId  string `json:"orgId"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
}

type AppResponse struct {
	Id      string `json:"id"`
	OrgId   string `json:"orgId"`
	Key     string `json:"key"`
	Secret  string `json:"secret"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

//
// @Summary 创建app
// @Tags    app
// @Accept  json
// @Produce  json
// @Param body body app.CreateAppRequest true "请求参数"
// @Success 200 {string} string "{"msg": "ok"}"
// @Failure 400 {string} json "{"error": "error info"}"
// @Failure 500 {string} json "{"error": "error info"}"
// @Router /app [post]
func Create(c *gin.Context) {
	r := &CreateAppRequest{}
	if err := util.ParseRequest(c, r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app := model.App{
		Id:      util.GetRandomString(11),
		OrgId:   r.OrgId,
		Key:     r.Key,
		Secret:  r.Secret,
		Name:    r.Name,
		Desc:    r.Desc,
		Created: time.Now(),
		Updated: time.Now(),
	}
	if err := appService.Create(&app); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"appId": app.Id})
}

//
// @Summary 获取app详情
// @Tags    app
// @Accept  json
// @Produce  json
// @Param id path string true "appId"
// @Success 200 {string} json "{"msg": "ok"}"
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /app/{id} [get]
func Get(c *gin.Context) {
	appId := c.Param("id")

	app, err := appService.Get(appId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02 15:04:05"
	appResponse := AppResponse{
		Id:      app.Id,
		OrgId:   app.OrgId,
		Key:     app.Key,
		Secret:  app.Secret,
		Name:    app.Name,
		Desc:    app.Desc,
		Created: app.Created.Format(layout),
		Updated: app.Updated.Format(layout),
	}

	c.JSON(http.StatusOK, appResponse)
}

//
// @Summary 获取app列表
// @Tags    app
// @Accept  json
// @Produce  json
// @Param orgId path string true "orgId"
// @Success 200 {string} json "{"msg": "ok"}"
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /apps/{orgId} [get]
func List(c *gin.Context) {
	orgId := c.Param("orgId")

	apps, err := appService.ListByOrgId(orgId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02 15:04:05"
	appListResponse := make([]AppResponse, 0)
	for _, a := range *apps {
		appListResponse = append(appListResponse, AppResponse{
			Id:      a.Id,
			OrgId:   a.OrgId,
			Key:     a.Key,
			Secret:  a.Secret,
			Name:    a.Name,
			Desc:    a.Desc,
			Created: a.Created.Format(layout),
			Updated: a.Updated.Format(layout),
		})
	}

	c.JSON(http.StatusOK, appListResponse)
}
