package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	dbSql, _ := db.DB()
	dbSql.SetMaxOpenConns(100)
	dbSql.SetMaxIdleConns(20)
	dbSql.SetConnMaxLifetime(30 * time.Second)
	DB = db
}
