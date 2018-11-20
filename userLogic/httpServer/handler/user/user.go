package user

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"msg/msgLogic/app"
	"msg/userLogic/service/model"
	userService "msg/userLogic/service/user"
	"msg/userLogic/util"
)

type RegByEmailRequest struct {
	Nick     string `json:"nick" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginByEmailRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Token   string    `json:"token"`
	UserId  string    `json:"userId"`
	Nick    string    `json:"nick"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Avt     string    `json:"avt"`
	Status  string    `json:"status"`
	Logined time.Time `json:"logined"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

//
// @Summary 注册by邮箱
// @Tags    用户
// @Accept  json
// @Produce  json
// @Param body body user.RegByEmailRequest true "请求参数"
// @Success 200 {string} string "{"msg": "ok"}"
// @Failure 400 {string} json "{"error": "error info"}"
// @Failure 500 {string} json "{"error": "error info"}"
// @Router /regByEmail [post]
func RegByEmail(c *gin.Context) {
	r := &RegByEmailRequest{}
	if err := util.ParseRequest(c, r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	passHash := util.EncryptPassword(r.Password, app.Config.Secret.PassHashSalt)
	now := time.Now()
	user := model.User{
		Id:       util.GetRandomString(11),
		Nick:     r.Nick,
		Email:    r.Email,
		Password: passHash,
		Created:  now,
		Updated:  now,
	}

	if err := userService.Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "reg success"})
}

//
// @Summary 登录by邮箱
// @Tags    用户
// @Accept  json
// @Produce  json
// @Param body body user.LoginByEmailRequest true "请求参数"
// @Success 200 {string} string "{"msg": "ok"}"
// @Failure 400 {string} json "{"error": "error info"}"
// @Failure 500 {string} json "{"error": "error info"}"
// @Router /loginByEmail [post]
func LoginByEmail(c *gin.Context) {
	r := &LoginByEmailRequest{}
	if err := util.ParseRequest(c, r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	whereMap := map[string]interface{}{
		"email": r.Email,
	}
	user, err := userService.Get(whereMap)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "cannot find email"})
	}

	token, err := util.BuildJwtToken(app.Config.Secret.JwtKey, jwt.MapClaims{
		"user_id": user.Id,
	})
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
	}

	userResponse := UserResponse{
		Token:   token,
		UserId:  user.Id,
		Nick:    user.Nick,
		Avt:     user.Avt,
		Status:  user.Status,
		Logined: user.Logined,
		Created: user.Created,
		Updated: user.Updated,
	}
	userService.UpdateLogined(user.Id, time.Now())

	c.JSON(http.StatusOK, userResponse)
}
