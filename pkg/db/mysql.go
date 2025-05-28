package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局数据库连接
var DB *gorm.DB

func InitMySQL(dsn string) (err error) {
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql connect fail:%v\n", err)
	}
	return
}

func CloseMySQL() error {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("mysql close fail:%v\n", err)
		return err
	}
	return sqlDB.Close()
}
