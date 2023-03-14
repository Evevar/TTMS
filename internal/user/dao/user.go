package dao

import (
	"TTMS/kitex_gen/user"
	"context"
)

func CreateUser(ctx context.Context, userInfo *user.User) error {
	return DB.WithContext(ctx).Create(userInfo).Error
}
