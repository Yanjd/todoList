package core

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"user/model"
	"user/services"
)

func BuildUser(user model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:       uint32(user.ID),
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
		UpdateAt: user.UpdatedAt.Unix(),
	}
	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) (err error) {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name=?", req.UserName).Debug().First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Code = 400
			resp.UserDetail = new(services.UserModel)
			return nil
		}
		resp.Code = 500
		resp.UserDetail = new(services.UserModel)
		return nil
	}

	if !user.CheckPassword(req.Password) {
		resp.Code = 400
		return nil
	}

	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) (err error) {
	resp.Code = 200
	if req.Password != req.PasswordConfirm {
		resp.Code = 400
		return errors.New("input not the same")
	}
	var tmp int64 = 0
	err = model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&tmp).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if tmp > 0 {
		return errors.New("username has exist")
	}
	user := model.User{
		UserName: req.UserName,
	}
	err = user.SetPassword(req.Password)
	if err != nil {
		resp.Code = 400
		return nil
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
