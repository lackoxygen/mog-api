package def

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/lackoxygen/mog-api/pkg/mgo"
)

var (
	x_database string = "x-database"
	x_collection string = "x-collection"
)

func Handle(c *gin.Context){
	var code int = 200
	var msg string
	database := c.Request.Header.Get(x_database)
	collection := c.Request.Header.Get(x_collection)
	if(database == ""){
		code = 500
		msg = "x_database is empty"
	}else if(collection == ""){
		code = 500
		msg = "x_collection is empty"
	}
	if code != 200 {
		c.JSON(http.StatusNotImplemented, gin.H{
			"code": code,
			"msg":  msg,
		})
		c.Abort()
		return
	}
	c.Set("d",database)
	c.Set("c",collection)
	c.Next()
}

func Inject(c *gin.Context)  {
	MgoSession := mgo.Session.Clone()

	defer MgoSession.Close()

	c.Set("mgo", MgoSession.DB(c.GetString("d")).C(c.GetString("c")))
	c.Next()
}