package models

import (
	"fmt"
	"gin-simple-project/global/variable"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	Id        int64  `gorm:"primaryKey" json:"id"`
	CreatedAt string `json:"created_at"` //日期时间字段统一设置为字符串即可
	UpdatedAt string `json:"updated_at"`
	//DeletedAt gorm.DeletedAt `json:"deleted_at"`   // 如果开发者需要使用软删除功能，打开本行注释掉的代码即可，同时需要在数据库的所有表增加字段deleted_at 类型为 datetime
}

func InitDB() {
	var dbDialector gorm.Dialector
	var err error
	dbDialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=false&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.database"),
		viper.GetString("mysql.charset")))
	variable.GormDbMysql, err = gorm.Open(dbDialector, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("数据库初始化失败 %v", err)
	}
	// 查询没有数据，屏蔽 gorm v2 包中会爆出的错误
	// https://github.com/go-gorm/gorm/issues/3789  此 issue 所反映的问题就是我们本次解决掉的
	_ = variable.GormDbMysql.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", MaskNotDataError)

	// https://github.com/go-gorm/gorm/issues/4838
	_ = variable.GormDbMysql.Callback().Create().Before("gorm:before_create").Register("CreateBeforeHook", CreateBeforeHook)
	// 为了完美支持gorm的一系列回调函数
	_ = variable.GormDbMysql.Callback().Update().Before("gorm:before_update").Register("UpdateBeforeHook", UpdateBeforeHook)

	if rawDb, err := variable.GormDbMysql.DB(); err != nil {
		log.Fatalf("数据库初始化失败 %v", err)
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(60 * time.Second)
		rawDb.SetMaxIdleConns(10)
		rawDb.SetMaxOpenConns(128)
	}
}
