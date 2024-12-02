package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PWD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	type Version struct {
		Version string
	}
	var version Version
	err = DB.Raw("select version() as version").Scan(&version).Error

	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{})
	fmt.Printf("mysql version: %s\n", version.Version)
}
