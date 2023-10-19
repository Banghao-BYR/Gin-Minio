package utils

import (
	"github.com/jinzhu/configor"
	"log"
)

var Config = struct {
	Minio struct {
		URL      string
		Username string
		Password string
	}
	Mysql struct {
		Host     string
		Username string
		Password string
		Dbname   string
		Config   string
	}
}{}

func LoadConf() {
	err := configor.Load(&Config, "config.yaml")
	if err != nil {
		log.Println("Init config error:", err.Error())
		return
	}
	log.Println("Init config:", Config)
}
