package action

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/lackoxygen/mog-api/model/input"
	"github.com/lackoxygen/mog-api/pkg/util"
	"fmt"
)

/**
删除
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