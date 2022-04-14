package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"mq-server/model"
	"strings"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string
)

func Init() {
	file, err := ini.Load("F:\\golandProjects\\TodoList\\todoList\\mq-server\\conf\\config.ini")
	if err != nil {
		fmt.Println("load ini config err: ", err)
		return
	}
	LoadMysql(file)
	LoadMQ(file)
	pathMySQL := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(pathMySQL)
	pathRabbitMQ := strings.Join([]string{RabbitMQ, "://", RabbitMQUser, ":", RabbitMQPassword, "@", RabbitMQHost, ":", RabbitMQPort, "/"}, "")
	model.RabbitMQ(pathRabbitMQ)
}

func LoadMysql(f *ini.File) {
	Db = f.Section("mysql").Key("Db").String()
	DbHost = f.Section("mysql").Key("DbHost").String()
	DbPort = f.Section("mysql").Key("DbPort").String()
	DbUser = f.Section("mysql").Key("DbUser").String()
	DbPassword = f.Section("mysql").Key("DbPassword").String()
	DbName = f.Section("mysql").Key("DbName").String()
}

func LoadMQ(f *ini.File) {
	RabbitMQ = f.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = f.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassword = f.Section("rabbitmq").Key("RabbitMQPassword").String()
	RabbitMQHost = f.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = f.Section("rabbitmq").Key("RabbitMQPort").String()
}
