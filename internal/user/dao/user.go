package dao

import (
	"TTMS/kitex_gen/user"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func getPassword(Password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("bcrypt.GenerateFromPassword error")
		return nil
	}
	return hash
}
func PasswordEqual(Password1, Password2 string) error {
	//hash, _ := bcrypt.GenerateFromPassword([]byte(Password1), bcrypt.DefaultCost)
	err := bcrypt.CompareHashAndPassword([]byte(Password2), []byte(Password1))
	if err == nil { //密码正确
		return nil
	}
	return err
}

// CreateUser 创建用户最后要加上判断 userType=9的条件
func CreateUser(ctx context.Context, userInfo *user.User) error {
	u := user.User{}
	if DB.Where("user_name = ?", userInfo.UserName).Find(&u); u.Id > 0 {
		return errors.New("该用户已存在")
	}
	userInfo.Password = string(getPassword(userInfo.Password))
	return DB.WithContext(ctx).Create(userInfo).Error
}

func UserLogin(ctx context.Context, userInfo *user.User) (user.User, error) {
	u := user.User{}
	DB.WithContext(ctx).Where("user_name = ?", userInfo.UserName).Find(&u)
	if u.Id > 0 {
		if PasswordEqual(userInfo.Password, u.Password) == nil {
			return u, nil
		} else {
			return u, errors.New("密码错误")
		}
	}
	return user.User{}, errors.New("未找到该用户")
}
func GetAllUser(ctx context.Context, current, pageSize int) ([]*user.User, error) {
	users := make([]*user.User, pageSize)
	tx := DB.WithContext(ctx).Offset((current - 1) * pageSize).Limit(pageSize).Find(&users)
	return users, tx.Error
}
func ChangeUserPassword(ctx context.Context, UserName string, Password, NewPassword string) error {
	u := user.User{}
	tx := DB.WithContext(ctx).Where("user_name = ?", UserName).Find(&u)
	if u.Id > 0 && PasswordEqual(Password, u.Password) == nil {
		u.Password = string(getPassword(NewPassword))
		tx = DB.Updates(&u)
	}
	return tx.Error
}

func DeleteUser(ctx context.Context, UserName string) error {
	return DB.WithContext(ctx).Where("user_name = ?", UserName).Delete(&user.User{}).Error
}
func GetUserInfo(ctx context.Context, UserName string) (*user.User, error) {
	u := user.User{}
	tx := DB.Where("user_name = ?", UserName).Find(&u)
	return &u, tx.Error
}
