package mysql

import (
	"gin_xiaomi_shop/models/admin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:ming199951@tcp(127.0.0.1:3306)/ginxiaomi?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, // 打印sql
	})
	if err != nil {
		return
	}
	err := DB.AutoMigrate(&admin.Manager{}, &admin.Role{}, &admin.Rights{})
	if err != nil {
		return
	}
}
