package input

/*import (
	"gopkg.in/mgo.v2/bson"
)
*/
type IdInput struct {
	Id	 string	`form:"id" binding:"required"`
}