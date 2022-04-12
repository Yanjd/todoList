package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"user/conf"
	"user/core"
	"user/services"
)

func main() {

	conf.Init()
	// etcd 注册件
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	// 微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),    //微服务名字
		micro.Address("127.0.0.1:9090"), //服务所在地址
		micro.Registry(etcdReg),         //etcd注册件
	)

	microService.Init()
	// 服务注册
	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 服务启动
	_ = microService.Run()
}
