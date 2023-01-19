package v1

import (
	"gin-simple-project/utils"
	"gin-simple-project/utils/response"

	"github.com/gin-gonic/gin"
)

func ApplicationCreate(c *gin.Context) {
	params := struct {
		SourceIp        string `json:"sourceIp" binding:"required,ip"`
		SourcePort      int    `json:"sourcePort" binding:"required,gte=1,lte=65535"`
		DestinationIp   string `json:"destinationIp" binding:"required,ip"`
		DestinationPort int    `json:"destinationPort" binding:"required,gte=1,lte=65535"`
	}{}
	valid, errs := utils.BindJsonAndValid(c, &params)
	if !valid {
		response.ErrorParamsValid(c, errs.Error(), nil)
		return
	}

	response.Success(c, "pong", nil)
}

func ApplicationList(c *gin.Context) {
	response.Success(c, "pong", nil)
}

func ApplicationDetail(c *gin.Context) {
	response.Success(c, "pong", nil)
}
func ApplicationUpdate(c *gin.Context) {
	response.Success(c, "pong", nil)
}
