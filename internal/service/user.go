package service

import (
	"errors"
	"live_room/internal/model"
	"live_room/pkg/db"
)

type User struct {
}

func (user *User) Login(username string, password string) (res *model.User, err error) {
	u := model.User{}
	res, err = u.FindUserByName(db.DB, username)
	if err != nil {
		return nil, err
	}
	if res.Password != password {
		return nil, errors.New("密码错误")
	}
	return res, nil
}
