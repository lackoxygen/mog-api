package main

import (
	"fmt"
	"log"
	"github.com/lackoxygen/mog-api/config"
	"github.com/lackoxygen/mog-api/router"
	"github.com/lackoxygen/mog-api/pkg/mgo"
)

func init() {
	config.Construct()
	mgo.Connect()
}

func main() {
	r := router.InitRouter()
	Addr := fmt.Sprintf(":%s", config.ServerConfig.Port)
	log.Printf("[info] start http server listening %s", Addr)
	r.Run(Addr)
}
