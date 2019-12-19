package util

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, msg string, data interface{})  {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":msg,
		"data":data,
	})
}

func Fail(c *gin.Context, msg string)  {
	c.JSON(500, gin.H{
		"code": 500,
		"msg":msg,
	})
}