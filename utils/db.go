package utils

import (
	"gin-minio/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDb() {
	// dsn := "root:Aa123456!@tcp(127.0.0.1:3306)/brain_controller?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := Config.Mysql.Username + ":" + Config.Mysql.Password + "@tcp(" + Config.Mysql.Host + ")/" + Config.Mysql.Dbname + "?" + Config.Mysql.Config
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Init db error:", err.Error())
		return
	}
	err = DB.AutoMigrate(&models.Image{})
	if err != nil {
		log.Println("Init db error:", err.Error())
		return
	}
	log.Println("Init db ...")
}
