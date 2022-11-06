package patient

import (
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/reservation"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/treatment_record"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/utils"
)

type Patient struct {
	utils.DbBaseModel
	FirstName        *string `gorm:"type:varchar(64)"`
	LastName         *string `gorm:"type:varchar(64)"`
	Age              *int    `gorm:"type:int4;default:0"`
	PhoneNumber      *string `gorm:"type:varchar(16)"`
	Email            *string `gorm:"type:varchar(64)"`
	AvatarUrl        *string `gorm:"type:text"`
	MedicalHistory   *string `gorm:"type:text"`
	Address          *string `gorm:"type:varchar(64)"`
	TreatmentRecords []*treatment_record.TreatmentRecord
	Reservations     []*reservation.Reservation
}

type Search struct {
	utils.DefaultSearchModel
}
