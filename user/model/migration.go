package model

import "fmt"

func migration() {
	DB = DB.Set("gorm:table_options", "charset=utf8mb4")
	err := DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("migration failed: ", err)
		return
	}
	//err := DB.AutoMigrate(&User{})
	//if err != nil {
	//	return
	//}
	//DB.Model()
}
