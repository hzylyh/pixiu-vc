package controller

import (
	"github.com/gin-gonic/gin"
	"pixiu-vc/utils"
)

func Demo(ctx *gin.Context) {
	utils.ResponseOk(ctx, "ok")
}
