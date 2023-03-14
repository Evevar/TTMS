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
	resp := &user.CreateUserResponse{}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	}
	return resp, err
}
