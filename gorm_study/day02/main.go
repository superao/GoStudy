package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义数据库模型
type User struct {
	gorm.Model
	Uuid  int64
	Name  string `gorm:"default:'小王子'"`
	Age   int32
	Hobby string
}

func main() {
	// 1. 连接数据库
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败！")
		return
	}
	// 2. 关闭数据库
	defer db.Close()

/*	// 建立自动迁移，当结构体中字段更新时，自动的将字段同步到数据表中
	db.AutoMigrate(&User{})
	u1 := &User{Uuid: 1000001, Name: "kukizhang", Age: 21, Hobby: "篮球"}

	result := db.Debug().Create(u1)         // Debug() 用于输出 SQL 语句
	fmt.Println(result.Error)				// 返回 error
	fmt.Println(result.RowsAffected)        // 返回插入记录的条数

	u2 := &User{Uuid: 1000002, Name: "ao", Age: 22, Hobby: ""}
	db.Debug().Select("Name", "Age", "Uuid").Create(u2)
	// INSERT INTO `users` (`uuid`,`name`,`age`) VALUES (1000002,'ao',22)

	u3 := &User{Uuid: 10000010, Name: "aoZhang", Age: 21, Hobby: "篮球"}
	db.Debug().Omit("Uuid", "Name").Create(u3)
	//INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`age`,`hobby`)
	//VALUES ('2021-06-14 17:38:23','2021-06-14 17:38:23',NULL,21,'篮球')
*/
	//u4 := User{Name: "hello123", Age: 18}
	//// 查询主键是否存在
	//fmt.Println(db.NewRecord(u4))
	//// 插入记录
	//db.Debug().Create(&u4)
	//// 查询主键是否存在
	//fmt.Println(db.NewRecord(u4))

	//// 向数据库中插入零值，检测是否会被忽略
	//u5 := User{Name: "", Age: 38}
	//db.Create(&u5)      // GORM 在创建记录的时候会自动忽略掉字段为零值的字段
	//// 所有字段的零值都不会保存在数据库中，在创建记录的时候这些字段会自动地被忽略掉

	// 插入时使用数据地默认值
	u6 := User{Age: 1200, Hobby: "123"}
	db.Debug().Create(&u6)




}
