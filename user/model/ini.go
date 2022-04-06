package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func Database(configString string) {
	db, err := gorm.Open(mysql.Open(configString), &gorm.Config{
		Logger: nil,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	dbSql, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSql.SetMaxOpenConns(100)
	dbSql.SetMaxIdleConns(20)
	dbSql.SetConnMaxLifetime(30 * time.Second)
	DB = db
	//migration()
}
