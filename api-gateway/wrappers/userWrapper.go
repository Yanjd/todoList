package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type userWrapper struct {
	client.Client
}

func (wrapper *userWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opt ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                30000,
		RequestVolumeThreshold: 20, //熔断器请求阈值
		ErrorPercentThreshold:  50,
		SleepWindow:            5000, //过多长时间，熔断器进行检测
	}

	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

func NewUserWrapper(c client.Client) client.Client {
	return &userWrapper{c}
}
