package conn

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建  mysql 数据库连接池
var DB *gorm.DB

func InitDB(user string, password string, host string, port string, dbname string) (err error) {
	fmt.Println("初始化数据库连接...")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		DB = nil
		return err
	}
	DB = db
	fmt.Println("数据库连接初始化完成")
	return nil

}
