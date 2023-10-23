package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/palp1tate/MultiLevelCommentDemo/global"
	"github.com/palp1tate/MultiLevelCommentDemo/model"
	"github.com/palp1tate/MultiLevelCommentDemo/router"
)

func InitMySQL() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"",
		"127.0.0.1",
		3306,
		"MultiLevelCommentDemo",
	)
	ormLogger := logger.Default
	global.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger, // 日志配置
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加s
		},
	})
	if err != nil {
		panic(err)
	}
	err = global.DB.AutoMigrate(
		model.User{},
		model.Moment{},
		model.Comment{},
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	InitMySQL()
	r := router.InitRouter()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
