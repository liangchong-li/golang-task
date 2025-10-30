package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
*/

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Posts    []Post
	Comments []Comment
}

// HashedPassword 密码加密
func (u *User) HashedPassword() (string, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPw), nil
}

// CheckPassword 校验密码
func (u *User) CheckPassword(pw string) error {
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
	return err
}
