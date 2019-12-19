package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lackoxygen/mog-api/config"
	"github.com/lackoxygen/mog-api/middleware/def"
	"github.com/lackoxygen/mog-api/controller/report"
	"github.com/lackoxygen/mog-api/controller/action"
	"github.com/lackoxygen/mog-api/controller/comment"
	"github.com/lackoxygen/mog-api/controller/sysmsg"
	"github.com/lackoxygen/mog-api/controller/payresult"
	"github.com/lackoxygen/mog-api/controller/review_attach"
)

func init()  {
	gin.SetMode(config.ServerConfig.Env)
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	r.Use(def.Handle)
	r.Use(def.Inject)

	//查询评论
	r.GET("/comment/list", comment.List)

	//添加评论
	r.POST("/comment/create", comment.Create)

	//查询系统消息分类
	r.GET("/sysmsg/class", sysmsg.GetSysMsgClass)

	//查询系统消息详情
	r.GET("/sysmsg/list", sysmsg.GetSysMsgClassList)

	//创建系统消息
	r.POST("/sysmsg/create", sysmsg.Create)

	//投诉用户
	r.POST("/report/create", report.Create)

	//删除table
	r.DELETE("/delete", action.Delete)

	//添加支付日志
	r.POST("/payresult/create", payresult.Create)

	r.POST("/review_attach/create", review_attach.Create)
	return r
}