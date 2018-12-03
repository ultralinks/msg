package link

import (
	"errors"
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

func Link(ctx *gin.Context) {
	timestamp := ctx.PostForm("time_stamp")
	clientSign := ctx.PostForm("sign")
	appKey := ctx.PostForm("app_key")
	userId := ctx.PostForm("user_id")
	nick := ctx.PostForm("nick")
	avt := ctx.PostForm("avt")
	action := ctx.PostForm("action")
	now := time.Now()

	// rpc调用获取secret
	app := rpc.FetchApp(appKey)
	if err := verify(timestamp, clientSign, appKey, app.Secret); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
	}

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

	if action == "linkToken" {
		ctx.JSON(http.StatusOK, gin.H{"token": lt.Token})
		return
	}

	if action == "link" {
		ctx.JSON(http.StatusOK, gin.H{
			"avt":  l.Avt,
			"nick": l.Nick,
			"key":  l.Key,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "action error"})
}

func verify(timestamp, clientSign, appKey, secret string) error {
	// 时间戳校验
	const FiveMinute = 300
	if timestamp == "" {
		return errors.New("timestamp can not empty")
	}
	clientNow, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return errors.New("timestamp error")
	}
	serverNow := time.Now().UTC().Unix()
	dValue := serverNow - clientNow
	if dValue < 0 {
		dValue = -dValue
	}
	if dValue > FiveMinute {
		return errors.New("timestamp error")
	}

	// 签名校验
	if clientSign == "" || appKey == "" {
		return errors.New("sign/app_key can not empty")
	}

	salt := "msg"
	serverSign := util.Md5(appKey + secret + salt)
	if clientSign != serverSign {
		return errors.New("sign error")
	}

	return nil
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
