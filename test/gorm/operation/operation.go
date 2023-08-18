package operation

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // gorm 自带字段
	Name       string
	Age        int
}

func Operation() {
	dsn := "root:root@tcp(192.168.60.120:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{CreateBatchSize: 100})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("连接MySQL成功")
	}
	// 自动迁移
	db.AutoMigrate(&User{})
	users := []User{
		{Name: "xiao li"},
		{Name: "xiao ming"},
	}
	db.CreateInBatches(users, 10)
}
