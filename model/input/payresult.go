package input

import (
	"gopkg.in/mgo.v2/bson"
)

type PayResultInput struct {
	Id 					bson.ObjectId		`bson:"_id"`
	Pay_method 	  		int32             	`json:"pay_method" binding:"required"`
	Out_trade_no		string 				`json:"out_trade_no" binding:"required"`
	Trade_no			string				`json:"trade_no" binding:"required"`
	Create_at			int32 				`json:"create_at" binding:"required"`
	Client_ip 			string 				`json:"client_ip" binding:"required"`
	Body 				map[string]interface{}
}
