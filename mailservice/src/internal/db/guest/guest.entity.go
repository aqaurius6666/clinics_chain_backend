package guest

import (
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/treatment_record"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/utils"
)

type Guest struct {
	utils.DbBaseModel
	FirstName        *string `gorm:"type:varchar(64)"`
	LastName         *string `gorm:"type:varchar(64)"`
	Age              *int    `gorm:"type:int4;default:0"`
	PhoneNumber      *string `gorm:"type:varchar(16)"`
	Email            *string `gorm:"type:varchar(64)"`
	MedicalHistory   *string `gorm:"type:text"`
	Address          *string `gorm:"type:varchar(64)"`
	TreatmentRecords []*treatment_record.TreatmentRecord
}

type Search struct {
	utils.DefaultSearchModel
}
