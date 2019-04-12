package app

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	mysqlArgs := Config.DB.User + ":" + Config.DB.Password + "@tcp(" + Config.DB.Host + ":" + Config.DB.Port + ")/" + Config.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	//Db, err = gorm.Open("mysql", "root:123456789@/chat?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("mysql", mysqlArgs)

	// 全局禁用表名复数
	DB.SingularTable(true)
	DB.DB().SetConnMaxLifetime(60 * time.Second)

	if err != nil {
		panic(err)
	}

	fmt.Println("DB connect success!!!", mysqlArgs)
	DB.LogMode(true)
}
