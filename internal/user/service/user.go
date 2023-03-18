package service

import (
	"TTMS/internal/user/dao"
	"TTMS/kitex_gen/user"
	"context"
	"log"
)

func CreateUserService(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	userInfo := &user.User{
		UserType: req.UserType,
		UserName: req.UserName,
		Password: req.Password,
	}
	err := dao.CreateUser(ctx, userInfo)
	resp := &user.CreateUserResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	log.Println("resp = ", resp)
	return resp, nil
}
func UserLoginService(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	userInfo := &user.User{
		UserName: req.UserName,
		Password: req.Password,
	}
	u, err := dao.UserLogin(ctx, userInfo)
	resp := &user.UserLoginResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
		resp.UserInfo = &u
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
		resp.UserInfo = &u
	}
	return resp, nil
}
func GetAllUserService(ctx context.Context, req *user.GetAllUserRequest) (*user.GetAllUserResponse, error) {
	resp := &user.GetAllUserResponse{BaseResp: &user.BaseResp{}}
	var err error
	resp.List, err = dao.GetAllUser(ctx, int(req.Current), int(req.PageSize))
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, err
}
func ChangeUserPasswordService(ctx context.Context, req *user.ChangeUserPasswordRequest) (*user.ChangeUserPasswordResponse, error) {
	err := dao.ChangeUserPassword(ctx, req.UserName, req.Password, req.NewPassword)
	resp := &user.ChangeUserPasswordResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, err
}
func DeleteUserService(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	err := dao.DeleteUser(ctx, req.UserName)
	resp := &user.DeleteUserResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, err
}
func GetUserInfoService(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	u, err := dao.GetUserInfo(ctx, req.UserName)
	resp := &user.GetUserInfoResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	resp.User = u
	return resp, err
}
