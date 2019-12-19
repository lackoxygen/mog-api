package payresult

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/lackoxygen/mog-api/model/input"
	"github.com/lackoxygen/mog-api/pkg/util"
)

/**
	create
 */
func Create(c *gin.Context)  {
	var body input.PayResultInput
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
