package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// 构造函数
func NewUser(username string) *User {
	return &User{
		Username:  username,
		CreatedAt: time.Now(),
	}
}

// 创建新用户
func (u *User) CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// 用id查询用户
func (u *User) FindUserById(db *gorm.DB, id int64) (*User, error) {
	user := &User{}
	if err := db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// 用name查询用户
func (u *User) FindUserByName(db *gorm.DB, username string) (*User, error) {
	user := &User{}
	if err := db.Where("username=?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) UpdateUserName(db *gorm.DB, id int64, newName string) error {
	return db.Model(&User{}).Where("id = ?", id).Update("username", newName).Error
}

func (u *User) DeleteUser(db *gorm.DB, id int64) error {
	return db.Delete(&User{}, id).Error
}
