package conv_link

import (
	"github.com/gin-gonic/gin"
	"msg/msgLogic/app"
	"msg/msgLogic/model"
	"net/http"
)

// @Summary Join Conv
// @Tags    Conv
// @Accept  json
// @Produce  json
// @Param body body conv.JoinConvRequest true "请求参数"
// @Success 200 {string} json "{"msg": "ok"}"
// @Router /conv[post]
func Create(c *gin.Context) {
	request := &JoinConvRequest{}
	c.ShouldBindJSON(&request)

	//ConvId, err := strconv.Atoi(request.ConvId)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, "")
	//}

	//fmt.Println("converid",ConvId)

	ConvLink := model.ConvLink{
		ConvId:    request.ConvId,
		UserToken: request.UserToken,
	}

	app.DB.Create(&ConvLink)

	c.JSON(http.StatusOK, "")
}
