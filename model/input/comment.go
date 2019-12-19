package input

import (
	"gopkg.in/mgo.v2/bson"
)

type CommentInput struct {
	Id 					bson.ObjectId		`bson:"_id"`
	User_id 	  		int32             	`json:"user_id" binding:"required"`
	To_user_id			int32 				`json:"to_user_id" binding:"required"`
	Activity_id			int32				`json:"activity_id" binding:"required"`
	Activity_type		int8 				`json:"activity_type"`
	Comment_user_id 	int32 				`json:"comment_user_id" binding:"required"`
	Comment 			string 				`json:"comment" binding:"required"`
	Create_at 			int32 				`json:"create_at" binding:"required"`
}