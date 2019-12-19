package mgo

import (
	"fmt"
	"github.com/lackoxygen/mog-api/config"
	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	Mongo *mgo.DialInfo
)

func GetMongoUrl() string{
	return fmt.Sprintf("mongodb://%s:%d/%s",
		config.MongodbConfig.Host,
		config.MongodbConfig.Port,
		config.MongodbConfig.Database,
	)
}

func Connect()  {
	uri := GetMongoUrl()
	mongo, err := mgo.ParseURL(uri)
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	Session = session
	Mongo = mongo
}
