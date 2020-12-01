package dao

import (
  "github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB *gorm.DB
)

func InitSqlite() (err error) {
	dsn := "/tmp/test.db"
	DB, err = gorm.Open("sqlite3", dsn)

	if err != nil {
		return
	}
	DB.SingularTable(true)          //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	DB.LogMode(true)                //打印sql语句
	//开启连接池
	DB.DB().SetMaxIdleConns(100)        //最大空闲连接
	DB.DB().SetMaxOpenConns(10000)      //最大连接数
	DB.DB().SetConnMaxLifetime(30)      //最大生存时间(s)
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}