package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"task/model"
	"task/service"
)

func (*TaskService) CreateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	ch, err := model.MQ.Channel()
	if err != nil {
		return fmt.Errorf("rabbitMQ channel err : %s", err.Error())
	}
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	body, _ := json.Marshal(req)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		return fmt.Errorf("rabbitMQ publish err : %s", err.Error())
	}
	return nil
}

func (*TaskService) GetTaskList(ctx context.Context, req *service.TaskRequest, resp *service.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var taskData []model.Task
	err := model.DB.Offset(int(req.Start)).Limit(int(req.Limit)).Where("uid = ?", req.Uid).Find(&taskData).Error
	if err != nil {
		return fmt.Errorf("mysql find err: %s", err.Error())
	}
	var count int64
	model.DB.Model(&model.Task{}).Where("uid = ?", req.Uid).Count(&count)
	var taskModels []*service.TaskModel
	for _, v := range taskData {
		taskModels = append(taskModels, BuildTask(v))
	}
	resp.TaskList = taskModels
	resp.Count = uint32(count)
	return nil
}

func (*TaskService) GetTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.First(&taskData, req.Id)
	taskRes := BuildTask(taskData)
	resp.TaskDetail = taskRes
	return nil
}

func (*TaskService) UpdateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.Model(&model.Task{}).Where("id = ?", req.Id).Where("uid = ?", req.Uid).First(&taskData)
	taskData.Title = req.Title
	taskData.Status = int(req.Status)
	taskData.Content = req.Content
	model.DB.Debug().Save(&taskData)
	resp.TaskDetail = BuildTask(taskData)
	return nil
}

func (*TaskService) DeleteTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	err := model.DB.Model(&model.Task{}).Where("id = ?", req.Id).Where("uid = ?", req.Uid).Debug().Delete(&model.Task{}).Error
	if err != nil {
		return fmt.Errorf("delete err: %s", err.Error())
	}
	return nil
}
