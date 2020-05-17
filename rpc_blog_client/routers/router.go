package routers

import (
	"github.com/gin-gonic/gin"
	"rpc_blog_client/conf"
	"rpc_blog_client/routers/api"
	"time"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	gin.SetMode(conf.Config.Server.RunMode)

	r.GET("/wait", func(context *gin.Context) {
		time.Sleep(10 * time.Second)
	})
	//r.Use(jwt.JWT())

	tags := r.Group("/tag")
	{
		tags.GET("/:id", api.GetTag)
		//tags.GET("", api.GetTags)
		//tags.POST("", api.AddTag)
		//tags.PUT("", api.EditTag)
		//tags.DELETE("/:id", api.DeleteTag)
	}


	return r
}
