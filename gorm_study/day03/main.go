package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义数据库模型
type User struct {
	ID   int64
	Name string `gorm:"default:'小王子'"`
	Age  int64
}

// 创建一条记录，使用零值对其进行初始化，两种方法：
// 使用指针方式实现零值存入数据库
// 使用Scanner/Valuer接口方式实现零值存入数据库

////// User 使用指针
//type User struct {
//	ID   int64
//	Name *string `gorm:"default:'小王子'"`
//	Age  int64
//}

// User 使用 Scanner/Valuer
//type User struct {
//	ID int64
//	Name sql.NullString `gorm:"default:'小王子'"` // sql.NullString 实现了Scanner/Valuer接口
//	Age  int64
//}

func main() {
	// 1. 连接数据库
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败！")
		return
	}
	// 2. 关闭数据库
	defer db.Close()

	db.AutoMigrate(&User{})

	//// 创建一条数据，Name 采用默认值
	//u1 := User{Age: 12}
	//db.Debug().Create(&u1)
	//
	//// 创建一条记录，Name 空值自动进行忽略，Name 采用默认值
	//u2 := User{Age: 15, Name: ""}
	//db.Debug().Create(&u2)

	// 创建一条记录，使用零值对其进行初始化，两种方法：
	// 使用指针方式实现零值存入数据库
	// 使用Scanner/Valuer接口方式实现零值存入数据库
	//u1 := User{Age: 12}    // Name 采用了默认值
	//db.Debug().Create(&u1)
	//
	//// 采用字符串的空值去创建一条数据库记录
	//u2 := User{Age: 100, Name: new(string)}
	//db.Debug().Create(&u2)

	//u1 := User{Age: 12}    // Name 采用了默认值
	//db.Debug().Create(&u1)
	//
	//// 采用字符串的空值去创建一条数据库记录
	//u2 := User{Age: 100, Name: sql.NullString{
	//	String: "",
	//	Valid:  true,		// value 是否有效
	//}}
	//db.Debug().Create(&u2)
	//
	//u1 := User{}
	//db.Create(&u1)

	//不支持批量插入数据
	//uList := []User{{Name: "123", Age: 12}, {Name: "456", Age: 15}}
	//db.Create(&uList)

	//// 一般查询
	//// 根据主键查询第一条记录
	//u := &User{}
	//db.First(u)
	//fmt.Println(u)
	//
	//// 获取数据表中的第一条数据
	//u1 := &User{}
	//db.Debug().Take(u1)
	//fmt.Println(u1)
	//
	//// 根据主键查询最后一条记录
	//u2 := &User{}
	//db.Debug().Last(u2)
	//fmt.Println(u2)
	//
	//// 查询所有的记录
	//u3 := &User{}
	//db.Debug().Find(u3)
	//fmt.Println(u3)
	//
	//// 查询指定的某条记录(仅当主键为整型时可用)
	//u4 := &User{}
	//db.Debug().First(u4, 10)

	// where 条件查询
	u := &User{}
	str := "小王子"
	db.Debug().Where("name = ?", str).Where("id = ?", 3).Find(u)

	fmt.Println(u)






	



}

