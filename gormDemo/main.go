package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
}

type UserRequest struct {
	// @inject_tag: json:"user_name" form:"user_name" uri:"user_name"
	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	// @inject_tag: json:"password" form:"password" uri:"password"
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	// @inject_tag: json:"password_confirm" form:"password_confirm" uri:"password_confirm"
	PasswordConfirm string `protobuf:"bytes,3,opt,name=PasswordConfirm,proto3" json:"PasswordConfirm,omitempty"`
}

func init() {
	dsn := "root:password@tcp(127.0.0.1:3306)/todolist_demo?charset=utf8&parseTime=true"
	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	dbSql, _ := database.DB()
	dbSql.SetMaxOpenConns(100)
	dbSql.SetMaxIdleConns(20)
	dbSql.SetConnMaxLifetime(30 * time.Second)
	Gorm = database
	migration()
}

var Gorm *gorm.DB

func main() {
	userReq1 := UserRequest{
		UserName: "yjd1997",
	}
	userReq2 := UserRequest{
		UserName: "yy",
	}
	Query1(userReq1)
	fmt.Println("----------------------")
	Query2(userReq2)
}

func migration() {
	Gorm.Set("gorm:table_options", "charset=utf8mb4")
	err := Gorm.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("migration failed: ", err)
		return
	}
}

func Query1(request UserRequest) {
	var users []User
	_ = Gorm.Find(&users)
	for _, v := range users {
		fmt.Println(v.UserName)
	}
	var data2 User
	if err := Gorm.Where("user_name = ?", request.UserName).First(&data2).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(data2.UserName)
}

func Query2(request UserRequest) {
	var users []User
	_ = Gorm.Find(&users)
	for _, v := range users {
		fmt.Println(v.UserName)
	}
	var data2 User
	if err := Gorm.Where("user_name = ?", request.UserName).First(&data2).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(data2.UserName)
}
