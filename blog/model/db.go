package model

import (
	"blog/utils"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db, err = gorm.Open(sqlite.Open(utils.DbPath), &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
})

func InitDb() {
	if err != nil {
		log.Fatalln("Failed to connect database! err: ", err.Error())
	}

	var sqlDB, err1 = db.DB()
	if err1 != nil {
		log.Fatalln("Failed to get *sql.DB, err: ", err1.Error())
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{}, &Tag{})

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//	sqlDB.Close()

}
