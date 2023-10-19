package routers

import (
	"gin-minio/controllers"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	//上传多文件
	r.POST("/UploadMultipleFiles", controllers.UploadMultipleFiles)
}
