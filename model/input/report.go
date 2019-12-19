package input

import (
	"gopkg.in/mgo.v2/bson"
)

type ReportInput struct {
	Id					bson.ObjectId		`bson:"_id"`
	User_id 	  		int32				`json:"user_id" binding:"required" bson:"user_id"`
	Report_user_id  	int32				`json:"report_user_id" binding:"required" bson:"report_user_id"`
	Report_type			int32				`json:"report_type" binding:"required" bson:"report_type"`
	Create_at			int32				`json:"create_at" binding:"required" bson:"create_at"`
	Operation_at 		int32				`json:"operation_at" bson:"operation_at"`
	Operation_id 		int32				`json:"operation_id" bson:"operation_id"`
	Operation_status 	int32				`json:"operation_status" bson:"operation_status"`
	Warning				int32 				`json:"warning" bson:"warning"`
	Body 				ReportBodyInput
}

type ReportBodyInput struct {
	Reason				string 				`json:"reason" bson:"reason"`
	Remark				string 				`json:"remark" bson:"remark"`
	Images				[]string			`json:"images" bson:"images"`
	Activity_id			int32				`json:"activity_id" bson:"activity_id"`
	Activity_type		int32 				`json:"activity_type" bson:"activity_type"`
}