package handlers

import (
	"api-gateway/pkg/util"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userReq services.UserRequest
	PanicUserError(c.Bind(&userReq))

	userService := c.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicUserError(err)
	c.JSON(http.StatusOK, gin.H{
		"data": userResp,
	})
}

func UserLogin(c *gin.Context) {
	var userReq services.UserRequest
	PanicUserError(c.Bind(&userReq))
	userService := c.Keys["userService"].(services.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicUserError(err)
	token, err := util.GenerateToken(uint(userResp.UserDetail.ID))
	PanicUserError(err)
	c.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"msg":  "success",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}
