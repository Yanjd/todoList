package wrappers

import (
	"api-gateway/services"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"strconv"
)

func NewTask(id uint64, name string) *services.TaskModel {
	return &services.TaskModel{
		Id:         id,
		Title:      name,
		Content:    "响应超时",
		StartTime:  1000,
		EndTime:    1000,
		Status:     0,
		CreateTime: 1000,
		UpdateTime: 1000,
	}
}

// DefaultTasks 降级
func DefaultTasks(resp interface{}) {
	models := make([]*services.TaskModel, 0)
	var i uint64
	for i = 0; i < 10; i++ {
		models = append(models, NewTask(i, "降级"+strconv.Itoa(20+int(i))))
	}
	result := resp.(*services.TaskListResponse)
	result.TaskList = models
}

type TaskWrapper struct {
	client.Client
}

func (wrapper *TaskWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opt ...client.CallOption) error {
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
		DefaultTasks(resp)
		return err
	})
}

func NewTaskWrapper(c client.Client) client.Client {
	return &TaskWrapper{
		c,
	}
}
