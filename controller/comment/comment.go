package comment

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/lackoxygen/mog-api/pkg/util"
	"github.com/lackoxygen/mog-api/model/input"
/*	"fmt"*/
)

/**
	list
 */
func List(c *gin.Context) {
	type Search struct {
		Self 		int32 	`form:"self"`	//是否是本人
		ActivityId	int32	`form:"activity_id" binding:"required"`	//活动ID
		UserId 		int32 	`form:"user_id" binding:"required"`		//用户ID
		Limit		int 	`form:"limit" binding:"required"`		//查询数量
	}
	var keyword Search
	err := c.ShouldBind(&keyword)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	var data []interface{}
	col := c.MustGet("mgo").(*mgo.Collection)
	condition := bson.M{}
	condition["activity_id"] = keyword.ActivityId
	if keyword.Self == 0 {
		condition["comment_user_id"] = keyword.UserId
	}
	col.Find(condition).Limit(keyword.Limit).All(&data)
	util.Success(c, "ok", data)
}

/**
	create
 */
func Create(c *gin.Context)  {
	var body input.CommentInput
	err := c.ShouldBind(&body)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	body.Id = bson.NewObjectId()
	col := c.MustGet("mgo").(*mgo.Collection)
	err = col.Insert(body)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	var response = make(map[string]interface{})
	response["insertId"] = body.Id
	util.Success(c, "success", response)
}
