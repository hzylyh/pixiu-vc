package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pixiu-vc/controller"
)

/**
路由控制
*/
func MapRoutes() *gin.Engine {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))

	// 之后根据具体路径配置静态资源目录
	router.LoadHTMLGlob("web/dist/*.html")          // 添加入口index.html
	router.LoadHTMLFiles("web/static/*/*")          // 添加资源路径
	router.Static("/static", "./web/dist/static")   // 添加资源路径
	router.StaticFile("/", "./web/dist/index.html") //前端接口

	// 创建根路由
	apiRoot := router.Group("/api")
	apiRoot.Use()

	// demo
	demoApi := apiRoot.Group("/demo")
	demoApi.POST("/demo1", controller.Demo)

	return router
}
