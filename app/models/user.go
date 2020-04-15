package models

import (
	"doc-system/config"

	"github.com/jinzhu/gorm"
)

// UserModel 用户model
type UserModel struct {
	gorm.Model
	Name     string `gorm:"column:name;not nul;index:name_idx"`
	Sex      int    `gorm:"column:sex;not null"` // 0 男  1 女
	Username string `gorm:"column:username;unique_index;not null"`
	Password string `gorm:"column:password;"`
	Token    string `gorm:"column:token"`
}

var self *gorm.DB = config.Self

// self.AutoMigrate(&UserModel{})

// TableName 用户表名
func (u *UserModel) TableName() string {
	return "users"
}

// Create 创建用户
func (u *UserModel) Create() error {
	return self.Create(&u).Error
}

// CheckUsernameExist 查看用户是否存在
func (u *UserModel) CheckUsernameExist(username string) bool {
	var user UserModel
	if self.Where("username = ?", username).First(&user).RecordNotFound() {
		return false
	}
	return true
}
