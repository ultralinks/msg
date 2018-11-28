package link

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"msg/msgLogic/rpc"
	linkService "msg/msgLogic/service/link"
	linkTokenService "msg/msgLogic/service/linkToken"
	"msg/msgLogic/service/model"
	"msg/msgLogic/util"
)

type CreateRequest struct {
	TimeStamp string `json:"time_stamp"`
	AppKey    string `json:"app_key"`
	Sign      string `json:"sign"`
	UserId    string `json:"user_id"`
	Nick      string `json:"nick"`
	Avt       string `json:"avt"`
}

//
// @Summary 创建link
// @Description 返回link_token
// @Tags    link
// @Accept  json
// @Produce  json
// @Param body body link.CreateRequest true "请求参数"
// @Success 200 {string} json "{"msg": "ok"}"
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /links [post]
func Create(ctx *gin.Context) {
	// 时间戳校验
	const FiveMinute = 300
	timestamp := ctx.PostForm("time_stamp")
	if timestamp == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": "timestamp can not empty"})
		return
	}
	clientNow, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "timestamp error"})
		return
	}
	serverNow := time.Now().UTC().Unix()
	dValue := serverNow - clientNow
	if dValue < 0 {
		dValue = -dValue
	}
	if dValue > FiveMinute {
		ctx.JSON(http.StatusOK, gin.H{"error": "timestamp error"})
		return
	}

	// 签名校验
	clientSign := ctx.PostForm("sign")
	appKey := ctx.PostForm("app_key")
	if clientSign == "" || appKey == "" {
		ctx.JSON(http.StatusOK, gin.H{"error": "sign/app_key can not empty"})
		return
	}
	// rpc调用获取secret
	app := rpc.FetchApp(appKey)
	salt := "msg"
	serverSign := util.Md5(appKey + app.Secret + salt)
	if clientSign != serverSign {
		ctx.JSON(http.StatusOK, gin.H{"error": "sign error"})
		return
	}

	userId := ctx.PostForm("user_id")
	nick := ctx.PostForm("nick")
	avt := ctx.PostForm("avt")
	now := time.Now()
	id := util.GetRandomString(11)
	link := model.Link{
		Id:      id,
		Nick:    nick,
		Avt:     avt,
		Key:     userId,
		AppId:   app.Id,
		Created: now,
		Updated: now,
	}
	linkToken := model.LinkToken{
		LinkId: id,
		Token:  util.Md5(id + userId + "msg"),
	}

	if err := linkService.Create(&link, &linkToken); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "database error"})
		return
	}

	l, _ := linkService.GetByKey(userId)
	lt, _ := linkTokenService.Get(l.Id)
	ctx.JSON(http.StatusOK, gin.H{"link_token": lt.Token})
}

//
// @Summary 通过linkKey获取link
// @Description 返回link
// @Tags    link
// @Accept  json
// @Produce  json
// @Param key path string true "linkKey"
// @Success 200 {string} json "{"msg": "ok"}"
// @Failure 400 {string} json "{"msg": "error info"}"
// @Failure 500 {string} json "{"msg": "error info"}"
// @Router /link/{key} [get]
func Get(ctx *gin.Context) {
	linkKey := ctx.Param("key")
	link, err := linkService.GetByKey(linkKey)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, link)
}
