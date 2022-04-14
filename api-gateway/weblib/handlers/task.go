package handlers

import (
	"api-gateway/pkg/util"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTaskList(c *gin.Context) {
	var taskReq services.TaskRequest
	PanicTaskError(c.Bind(&taskReq))

	taskService := c.Keys["taskService"].(services.TaskService)

	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)

	taskListResp, err := taskService.GetTaskList(context.Background(), &taskReq)
	PanicTaskError(err)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"tasks": taskListResp.TaskList,
			"count": taskListResp.Count,
		},
	})
}

func GetTask(c *gin.Context) {
	var taskReq services.TaskRequest
	PanicTaskError(c.Bind(&taskReq))

	taskService := c.Keys["taskService"].(services.TaskService)

	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	taskReq.Id = id

	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)

	taskResp, err := taskService.GetTask(context.Background(), &taskReq)
	PanicTaskError(err)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "get task success",
		"data": taskResp,
	})
}

func CreateTask(c *gin.Context) {
	var taskReq services.TaskRequest
	PanicTaskError(c.Bind(&taskReq))

	taskService := c.Keys["taskService"].(services.TaskService)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)

	taskResp, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicTaskError(err)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "create task success",
		"data": taskResp,
	})
}

func UpdateTask(c *gin.Context) {
	var taskReq services.TaskRequest
	PanicTaskError(c.Bind(&taskReq))

	taskService := c.Keys["taskService"].(services.TaskService)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	taskReq.Id = id

	taskResp, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicTaskError(err)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "update task success",
		"data": taskResp,
	})
}

func DeleteTask(c *gin.Context) {
	var taskReq services.TaskRequest
	PanicTaskError(c.Bind(&taskReq))

	taskService := c.Keys["taskService"].(services.TaskService)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	taskReq.Id = id

	taskResp, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicTaskError(err)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "delete task success",
		"data": taskResp,
	})
}
