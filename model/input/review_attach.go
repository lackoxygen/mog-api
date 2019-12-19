package input

import (
	"gopkg.in/mgo.v2/bson"
)

type ReviewAttachInput struct {
	Id					bson.ObjectId		`bson:"_id"`
	User_id 	  		int32				`json:"user_id" binding:"required" bson:"user_id"`
	Review_type  		int32				`json:"review_type" bson:"review_type"`
	Review_status		int32				`json:"review_status" bson:"review_status"`
	Review_at			int32				`json:"review_at" bson:"review_at"`
	Media_type 			int32				`json:"media_type" bson:"media_type"`
	Note 				string				`json:"note" bson:"note"`
	Sex 				int32				`json:"sex" bson:"sex"`
	Operator_id			int32 				`json:"operator_id" bson:"operator_id"`
	Create_at			int32 				`json:"create_at" binding:"required" bson:"create_at"`
	Body 				ReviewAttachBodyInput
}

type ReviewAttachBodyInput struct {
	Nymph_id			int32			`json:"nymph_id" bson:"nymph_id"`
	Real_img			string 			`json:"real_img" bson:"real_img"`
	Face_img			string 			`json:"face_img" bson:"face_img"`
	Picture				string 			`json:"picture" bson:"picture"`
	Media_type			int32 			`json:"media_type" bson:"media_type"`
	Picture_id			int32 			`json:"picture_id" bson:"picture_id"`
	Match_score			int32			`json:"match_score" bson:"match_score"`
}