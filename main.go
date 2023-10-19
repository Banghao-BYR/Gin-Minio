package main

import (
	"gin-minio/routers"
	"gin-minio/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	r := gin.Default()
	routers.RouterInit(r)
	r.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
}

func Init() {
	utils.LoadConf()
	utils.InitDb()
	utils.InitMinio()
}
