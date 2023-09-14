package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func DataBase(connstring string) {
	fmt.Println(connstring)
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		panic("Mysql数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)                       //设置表名不为复数
	db.DB().SetMaxIdleConns(20)                  //??好像是设置连接池，看到？后查资料看啥意思
	db.DB().SetMaxOpenConns(200)                 //设置最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) //设置连接时间？？似乎是最大生命周期
	DB = db
	migration()
}
