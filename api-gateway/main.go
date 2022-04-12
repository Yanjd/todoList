package main

import (
	"api-gateway/services"
	"api-gateway/weblib"
	"api-gateway/wrappers"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {
	etcdR := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)

	// 用户服务调用实例
	userService := services.NewUserService("rpcUserService", userMicroService.Client())
	// 创建微服务实例， 使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address(":4000"),
		web.Handler(weblib.NewRouter(userService)),
		web.Registry(etcdR),
		web.RegisterTTL(30*time.Second),
		web.RegisterInterval(15*time.Second),
		web.Metadata(map[string]string{
			"protocol": "http",
		}),
	)
	_ = server.Init()
	_ = server.Run()
}
