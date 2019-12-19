package input

import (
	"gopkg.in/mgo.v2/bson"
)

type SysmsgInput struct {
	Id					bson.ObjectId		`bson:"_id"`
	User_id 	  		int32				`json:"user_id" binding:"required" bson:"user_id"`
	Ry_send_user_id 	string				`json:"ry_send_user_id" binding:"required" bson:"ry_send_user_id"`
	All_recv			int32 				`json:"all_recv" bson:"all_recv"`
	To_user_id			int32 				`json:"to_user_id" bson:"to_user_id"`
	Ry_recv_user_id		string 				`json:"ry_recv_user_id" binding:"required" bson:"ry_recv_user_id"`
	Msg_title			string 				`json:"msg_title" binding:"required" bson:"msg_title"`
	Msg_time			string 				`json:"msg_time" binding:"required" bson:"msg_time"`
	Msg_type			string 				`json:"msg_type" binding:"required" bson:"msg_type"`
	Msg_type_id			int32 				`json:"msg_type_id" binding:"required" bson:"msg_type_id"`
	Body				SysmsgBodyInput
}

type SysmsgBodyInput struct {
	Activity_id			int32 				`json:"activity_id" bson:"activity_id"`
	Activity_type		int32 				`json:"activity_type" bson:"activity_type"`
	Form_user_id		int32				`json:"form_user_id" bson:"form_user_id"`
	Picture_id			int32				`json:"picture_id" bson:"picture_id"`
	Read_pack_type		int32				`json:"read_pack_type" bson:"read_pack_type"`
	Image				string 				`json:"image" bson:"image"`
	Is_burn				int32				`json:"is_burn" bson:"is_burn"`
	Status				int32 				`json:"status" bson:"status"`
	Album_confirm_id	int32				`json:"album_confirm_id" bson:"album_confirm_id"`
	Aocial_account		string 				`json:"aocial_account" bson:"aocial_account"`
	Lock_user_id		int32				`json:"lock_user_id" bson:"lock_user_id"`
}