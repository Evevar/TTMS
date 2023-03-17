package dao

import (
	"TTMS/kitex_gen/user"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, userInfo *user.User) error {
	u := user.User{}
	if DB.Where("user_name = ?", userInfo.UserName).Find(&u); u.Id > 0 {
		return errors.New("该用户已存在")
	}
	return DB.WithContext(ctx).Create(userInfo).Error
}

func UserLogin(ctx context.Context, userInfo *user.User) (user.User, error) {
	u := user.User{}
	tx := DB.Where("user_name = ?", userInfo.UserName).Find(&u)
	if u.Id > 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
		if bcrypt.CompareHashAndPassword([]byte(u.Password), hash) == nil { //密码正确
			return u, nil
		}
	}
	return user.User{}, tx.Error
}
func GetAllUser(ctx context.Context, current, pageSize int) ([]*user.User, error) {
	users := make([]*user.User, pageSize)
	tx := DB.Offset((current - 1) * pageSize).Limit(pageSize).Find(&users)
	return users, tx.Error
}
