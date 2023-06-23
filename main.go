package main

import (
	"gin-minio/api"
	"gin-minio/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	RouterInit(r)
	utils.InitMinio()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func RouterInit(r *gin.Engine) {
	//上传多文件
	r.POST("/UploadMultipleFiles", api.UploadMultipleFiles)
}
