package sysmsg

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/lackoxygen/mog-api/model/input"
	"github.com/lackoxygen/mog-api/pkg/util"
)
/**
	查询消息通
 */
func GetSysMsgClass(c *gin.Context)  {
	type Search struct {
		UserId	int32	`form:"user_id" binding:"required"`
	}
	var keyword Search
	err := c.ShouldBind(&keyword)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	fmt.Println(keyword)
	col := c.MustGet("mgo").(*mgo.Collection)

	match := bson.M{"$match": bson.M{"to_user_id": keyword.UserId}}
	group := bson.M{"$group": bson.M{"_id":"$msg_type_id","msg_title":bson.M{"$last":"$msg_title"},"msg_time":bson.M{"$last":"$msg_time"},"read":bson.M{"$last":"$read"}}}
	sort := bson.M{"$sort": bson.M{"msg_time": -1}}
	pipeline := []bson.M{
		match,
		group,
		sort,
	}
	var data []interface{}

	pipe := col.Pipe(pipeline)
	err = pipe.All(&data)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	util.Success(c, "ok", data)
}

func GetSysMsgClassList(c *gin.Context)  {
	type Param struct {
		MsgTypeId	int32	`form:"msg_type_id" binding:"required"`
		UserId		int32 	`form:"user_id" binding:"required"`
	}
	var paginate input.Paginate
	var param Param
	err1 := c.ShouldBind(&paginate)
	err2 := c.ShouldBind(&param)
	if err1 != nil || err2 != nil{
		var msg string
		if (err1 != nil) {
			msg = err1.Error()
		}else{
			msg = err2.Error()
		}
		util.Fail(c, msg)
		return
	}
	var data []interface{}

	var skip int = (paginate.Page - 1) * paginate.Limit

	col := c.MustGet("mgo").(*mgo.Collection)
	err := col.Find(bson.M{"to_user_id":param.UserId, "msg_type_id":param.MsgTypeId}).Skip(skip).Limit(paginate.Limit).All(&data)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	util.Success(c, "ok", data)
}

func Create(c *gin.Context)  {
	var body input.ReviewAttachInput
	err := c.ShouldBind(&body)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	lastId := bson.NewObjectId()
	body.Id = lastId
	col := c.MustGet("mgo").(*mgo.Collection)
	err = col.Insert(body)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	var response = make(map[string]interface{})
	response["insertId"] = lastId
	util.Success(c, "success", response)
}