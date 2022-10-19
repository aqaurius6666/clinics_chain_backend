package user

import (
	"github.com/minh1611/go_structure/authservice/src/internal/utils"
)

type User struct {
	utils.DbBaseModel
	Name        *string `gorm:"type:varchar(128)"`
	Age         *int16  `gorm:"type:int4;default:0"`
	PhoneNumber *string `gorm:"type:varchar(16)"`
	Email       *string `gorm:"type:varchar(64)"`
	Password    *string `gorm:"type:varchar(64)"`
}

type Search struct {
	utils.DbDefaultSearchModel
}