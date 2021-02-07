package service

import "github.com/gin-gonic/gin"

var DemoSrv = &demoService{}

type demoService struct {
}

func (srv *demoService) Demo1(c *gin.Context) string {
	return "demo"
}
