package models

import "gorm.io/gorm"

type Image struct {
	//id,bucket_name ,folder_name,url,add_time,update_time,is_delete
	gorm.Model
	BucketName string
	FolderName string
	Url        string
	IsDelete   *bool `gorm:"default:false"`
}
