package service_type

import "github.com/minh1611/clinics_chain_management/mailservice/src/internal/utils"

type ServiceType struct {
	utils.DbBaseModel
	ServiceName *string  `gorm:"type:text"`
	Fee         *float64 `gorm:"type:float"`
}

type Search struct {
	utils.DefaultSearchModel
}
