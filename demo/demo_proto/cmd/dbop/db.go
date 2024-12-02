package main

import (
	// "fmt"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()

	//CURD

	// mysql.DB.Create(&model.User{Mail: "3473901836@qq.com", Password: "123456"})

	// mysql.DB.Model(&model.User{}).Where("mail = ?", "3473901836@qq.com").Update("password", "654321")

	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("mail = ?", "3473901836@qq.com").First(&row)
	// fmt.Printf("row : %+v\n", row);

	//软删除，
	//默认情况下，GORM 的 Delete 方法会执行软删除（即设置 deleted_at 字段为当前时间，
	//而不是真正从数据库中删除记录）。如果表结构中没有 deleted_at 字段，则会真正删除记录。
	// mysql.DB.Where("mail = ?", "3473901836@qq.com").Delete(&model.User{})

	// 硬删除 使用 Unscoped() 方法后，GORM 会忽略软删除机制，真正从数据库中删除记录，无论这些记录是否已经被软删除。
	mysql.DB.Unscoped().Where("mail = ?", "3473901836@qq.com").Delete(&model.User{})
}
