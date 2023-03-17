package service

import (
	"TTMS/internal/user/dao"
	"TTMS/kitex_gen/user"
	"context"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateUserService(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("bcrypt.GenerateFromPassword error")
		return nil, err
	}
	userInfo := &user.User{
		UserType: req.UserType,
		UserName: req.UserName,
		Password: string(hash),
	}
	err = dao.CreateUser(ctx, userInfo)
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
