package main

import (
	"TTMS/internal/user/service"
	user "TTMS/kitex_gen/user"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	//获取resp
	resp, err = service.CreateUserService(ctx, req)
	return resp, err
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...

	return service.UserLoginService(ctx, req)
}

// GetAllUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetAllUser(ctx context.Context, req *user.GetAllUserRequest) (resp *user.GetAllUserResponse, err error) {
	// TODO: Your code here...

	return service.GetAllUserService(ctx, req)
}

// ChangeUserPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserPassword(ctx context.Context, req *user.ChangeUserPasswordRequest) (resp *user.ChangeUserPasswordResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}
