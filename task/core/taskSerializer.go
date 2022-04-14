package core

import (
	"task/model"
	"task/service"
)

func BuildTask(val model.Task) *service.TaskModel {
	return &service.TaskModel{
		Id:         uint64(val.ID),
		Uid:        uint64(val.Uid),
		Title:      val.Title,
		Content:    val.Content,
		StartTime:  val.StartTime,
		EndTime:    val.EndTime,
		Status:     int64(val.Status),
		CreateTime: val.CreatedAt.Unix(),
		UpdateTime: val.UpdatedAt.Unix(),
	}
}
