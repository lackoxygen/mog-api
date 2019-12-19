package report

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/lackoxygen/mog-api/model/input"
	"github.com/lackoxygen/mog-api/pkg/util"
	"fmt"
)

func init() {
	fmt.Println("-------report---------")
}

/**
	创建投诉
 */
func Create(c *gin.Context)  {
	var body input.ReportInput
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


/**
	删除投诉
 */
func Delete(c *gin.Context) {
	var param input.IdInput
	err := c.ShouldBind(&param)
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	fmt.Println(param)
	col := c.MustGet("mgo").(*mgo.Collection)
	err = col.RemoveId(bson.ObjectIdHex(param.Id))
	if err != nil {
		util.Fail(c, err.Error())
		return
	}
	var response = make(map[string]interface{})
	response["deleteId"] = param.Id
	util.Success(c, "success", response)
}