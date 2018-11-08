package conv

import (
	"github.com/gin-gonic/gin"
	"msg/msgLogic/app"
	"msg/msgLogic/model"
	"net/http"
)

// @Summary Create Conv
// @Tags    Conv
// @Accept  json
// @Produce  json
// @Param body body conv.CreateConvRequest true "请求参数"
// @Success 200 {string} json "{"msg": "ok"}"
// @Router /conv[post]
func Create(c *gin.Context) {
	request := &CreateConvRequest{}
	c.ShouldBindJSON(&request)

	Conv := model.Conv{
		CreateUserToken: request.Token,
		Name:            request.Name,
	}

	app.DB.Create(&Conv)

	ConvLink := model.ConvLink{
		ConvId:    Conv.Id,
		UserToken: Conv.CreateUserToken,
	}

	app.DB.Create(&ConvLink)

	c.JSON(http.StatusOK, Conv)
}
