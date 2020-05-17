package app

import (
	"github.com/gin-gonic/gin"
	"rpc_blog_client/pkg/e"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type Gin struct {
	c *gin.Context
}

func GetResponse(code int, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	}
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.c.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}
