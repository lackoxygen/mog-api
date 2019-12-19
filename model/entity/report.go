package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type Report struct {
	Id     				bson.ObjectId 		`json:"id" bson:"_id,omitempty"`
	UserId 	  			int32             	`json:"user_id" bson:"user_id"`
	ReportUserId  		int32             	`json:"report_user_id" bson:"report_user_id"`
	ReportType			int32				`json:"report_type" bson:"report_type"`
	CreateAt			int32 				`json:"create_at" bson:"create_at"`
	OperationAt 		int32 				`json:"operation_at" bson:"operation_at"`
	OperationId 		int32 				`json:"operation_id" bson:"operation_id"`
	OperationStatus 	int32 				`json:"operation_status" bson:"operation_status"`
	Warning				int32 				`json:"warning" bson:"warning"`
	Body				Body
}

type Body struct {
	ActivityId			int32 				`json:"activity_id" bson:"activity_id"`
	Remark				string 				`json:"remark" bson:"remark"`
	ActivityType 		int32 				`json:"activity_type" bson:"activity_type"`
	Reason				string 				`json:"reason" bson:"reason"`
	Images				[]string			`json:"images" bson:"images"`
}