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
func CreateUser(ctx context.Context, userInfo *user.User) (err error) {
	u := user.User{}
	if DB.Where("user_name = ?", userInfo.UserName).Find(&u); u.Id > 0 {
		err = errors.New("警告，您设置了重名用户")
	}
	userInfo.Password = string(getPassword(userInfo.Password))
	DB.WithContext(ctx).Create(userInfo)
	return err
}

func UserLogin(ctx context.Context, userInfo *user.User) (user.User, error) {
	u := user.User{}
	DB.WithContext(ctx).Where("id = ?", userInfo.Id).Find(&u)
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
func ChangeUserPassword(ctx context.Context, UserId int, Password, NewPassword string) error {
	u := user.User{}
	tx := DB.WithContext(ctx).Where("id = ?", UserId).Find(&u)
	if u.Id > 0 && PasswordEqual(Password, u.Password) == nil {
		u.Password = string(getPassword(NewPassword))
		tx = DB.WithContext(ctx).Updates(&u)
	}
	return tx.Error
}

func DeleteUser(ctx context.Context, UserId int) error {
	return DB.WithContext(ctx).Where("id = ?", UserId).Delete(&user.User{}).Error
}
func GetUserInfo(ctx context.Context, UserId int) (*user.User, error) {
	u := user.User{}
	tx := DB.WithContext(ctx).Where("id = ?", UserId).Find(&u)
	return &u, tx.Error
}

func BindEmail(ctx context.Context, id int64, email string) error {
	u := user.User{}
	DB.WithContext(ctx).Select("where email = ?", email).Find(&u)
	if u.Id == 0 { //该邮箱尚未被其他用户绑定
		return DB.WithContext(ctx).Updates(&user.User{Id: id, Email: email}).Error
	} else if u.Id != id {
		return errors.New("该邮箱已经被其他用户绑定")
	}
	return nil
}
func ForgetPassword(ctx context.Context, Email string, NewPassword string) error {
	u := user.User{}
	DB.WithContext(ctx).Where("email = ?", Email).Find(&u)
	if u.Id > 0 {
		u.Password = string(getPassword(NewPassword))
		return DB.WithContext(ctx).Updates(&u).Error
	}
	return errors.New("该邮箱尚未绑定")
}
