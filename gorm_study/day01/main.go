package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 连接 MySql 数据库
func main() {
	// 1. 连接 MySQL 数据库
	// args = user:password@tcp(ip:port)/dbname?charset=utf8mb4&parseTime=true&loc=Local
	mysqlInfo := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open("mysql", mysqlInfo)
	if err != nil {
		fmt.Println("mysql connect error!")
		return
	}

	// 2. 设置数据库连接池
	sqlPool := db.DB()
	sqlPool.SetMaxIdleConns(10)    		// 设置空闲连接池中连接的最大数量
	sqlPool.SetConnMaxIdleTime(time.Minute) // 设置空闲连接的最大存活时间
	sqlPool.SetMaxOpenConns(100)   		// 设置打开数据库连接的最大数量
	sqlPool.SetConnMaxLifetime(time.Hour)	// 设置连接可复用的最大时间

	// 3. 关闭数据库连接
	defer db.Close()
}
