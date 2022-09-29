package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        *string `gorm:"type:varchar(128)"`
	Age         *int16  `gorm:"type:int4;default:0"`
	PhoneNumber *string `gorm:"type:varchar(16)"`
	Email       *string `gorm:"type:varchar(64)"`
	Password    *string `gorm:"type:varchar(64)"`
}
