package dal

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `json:"username" gorm:"uniqueIndex;not null"`
	Password    string `json:"password" gorm:"not null"`
	PhoneNumber string `json:"phonenumber" gorm:"uniqueIndex;not null"`
}

func CreateUser(ctx context.Context, user *User) error {
	if err := DB.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func QueryUser(ctx context.Context, phonenumberOrUserName string) (*User, error) {
	var res *User
	if err := DB.WithContext(ctx).Model(&User{}).Where("username = ?", phonenumberOrUserName).Find(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
