package routers

import (
	v1 "gin-simple-project/controller/v1"
	"gin-simple-project/middleware"
	"gin-simple-project/utils/response"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations("zh"))
	r.Use(middleware.AddTraceId()) //添加全局的traceId

	apiv1 := r.Group("/v1/example")
	{
		apiv1.GET("/ping", func(c *gin.Context) {
			response.Success(c, "pong", nil)
		})
		applicationRouters := apiv1.Group("/dyeing")
		applicationRouters.Use()
		{
			applicationRouters.GET("/application", v1.ApplicationList)       //申请列表
			applicationRouters.POST("/application", v1.ApplicationCreate)    //申请
			applicationRouters.GET("/application/:id", v1.ApplicationDetail) //申请详情
			applicationRouters.PUT("/application:id", v1.ApplicationUpdate)  //更新申请
		}
	}

	openApiv1 := r.Group("/open/v1/example")
	{
		openApiv1.GET("/ping", func(c *gin.Context) {
			response.Success(c, "pong", nil)
		})
	}

	return r
}
