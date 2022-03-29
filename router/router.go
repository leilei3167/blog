package router

import (
	v1 "blog/router/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	//后续扩充中间件

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)

		//在修改时,考虑如何修改?根据id的话可以通过路径参数传入
		apiv1.PUT("/tags/:id", v1.UpdateTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

	}

	return r
}
