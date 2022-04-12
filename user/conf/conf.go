package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"user/model"
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
	file, err := ini.Load("F:\\golandProjects\\TodoList\\todoList\\user\\conf\\config.ini")
	if err != nil {
		fmt.Println("load ini config err: ", err)
		return
	}
	LoadMysqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
}

func LoadMysqlData(f *ini.File) {
	Db = f.Section("mysql").Key("Db").String()
	DbHost = f.Section("mysql").Key("DbHost").String()
	DbPort = f.Section("mysql").Key("DbPort").String()
	DbUser = f.Section("mysql").Key("DbUser").String()
	DbPassword = f.Section("mysql").Key("DbPassword").String()
	DbName = f.Section("mysql").Key("DbName").String()
}
