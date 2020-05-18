package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"rpc_blog_client/pkg/app"
	"rpc_blog_client/pkg/e"
	"rpc_blog_client/rpc"
)

func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID不合法")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		c.JSON(http.StatusOK, app.GetResponse(e.INVALID_PARAMS, nil))
		return
	}

	data, code := rpc.GetArticle(id)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, app.GetResponse(code, nil))
		return
	}
	c.JSON(http.StatusOK, app.GetResponse(code, data))
}
