package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/user/conf"
	"github.com/cloudwego/kitex/pkg/klog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(
		conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PWD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Fatalf("Failed to connect to database: %v", err)
	}
	// 使用tracing.NewPlugin()为数据库连接添加追踪插件，如果添加失败则抛出错误
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	type Version struct {
		Version string
	}
	var version Version
	err = DB.Raw("select version() as version").Scan(&version).Error
	if err != nil {
		klog.Fatalf("Failed to execute SQL query: %v", err)
	}
	DB.AutoMigrate(&model.User{})
	fmt.Printf("mysql version: %s\n", version.Version)
}
