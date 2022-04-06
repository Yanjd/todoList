package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

//Db = mysql
//DbHost = 127.0.0.1
//DbPort = 3306
//DbUser = root
//DbPassword = password
//DbName = todolist_demo
var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("load ini config err: ", err)
		return
	}
	LoadMysqlData(file)
}

func LoadMysqlData(f *ini.File) {
	Db = f.Section("mysql").Key("Db").String()
	DbHost = f.Section("mysql").Key("DbHost").String()
	DbPort = f.Section("mysql").Key("DbPort").String()
	DbUser = f.Section("mysql").Key("DbUser").String()
	DbPassword = f.Section("mysql").Key("DbPassword").String()
	DbName = f.Section("mysql").Key("DbName").String()
}
