package variable

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DateFormat  = "2006-01-02 15:04:05" //  设置全局日期时间格式
	GormDbMysql *gorm.DB                // 全局gorm的客户端连接
	Rdb         *redis.Client
)
