package dal

import "github.com/cloudwego/biz-demo/gomall/app/cart/biz/dal/redis"

func Init() {
	redis.Init()
	// mysql.Init()
}
