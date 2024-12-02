package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mail     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(100);not null"`
}

// 自定义名字
func (User) TableName() string {
	return "user"
}
