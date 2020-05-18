package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"rpc_blog_client/conf"
)

func GetPage(c *gin.Context) (int, int) {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	size, _ := com.StrTo(c.Query("size")).Int()
	if size == 0 {
		size = conf.Config.App.PageSize
	}
	if page > 0 {
		result = (page - 1) * size
	}
	return result, size
}
