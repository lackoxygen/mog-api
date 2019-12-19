package config

import (
	"log"
	"github.com/go-ini/ini"
	"fmt"
)

type App struct {

}

type Server struct {
	Port string
	Env string
}

type Mongodb struct {
	Host string
	Port int
	User string
	Password string
	Database string
	Prefix string
}

var ServerConfig, MongodbConfig = &Server{}, &Mongodb{}

var cfg *ini.File

var configFile string = "config/config.ini"

func Construct()  {
	var err error
	cfg, err = ini.Load(configFile)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse '%s': %v", configFile, err)
	}
	mapTo("server", ServerConfig)
	mapTo("mongodb", MongodbConfig)
	fmt.Println(ServerConfig)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}