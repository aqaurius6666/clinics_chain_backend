package treatment_record

import (
	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/db/service_type"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/utils"
)

type TreatmentRecord struct {
	utils.DbBaseModel
	PatientID    uuid.UUID                   `gorm:"type:uuid"`
	GuestID      uuid.UUID                   `gorm:"type:uuid"`
	DoctorID     uuid.UUID                   `gorm:"type:uuid"`
	NurseID      uuid.UUID                   `gorm:"type:uuid"`
	ServiceTypes []*service_type.ServiceType `gorm:"many2many:record_serviceTypes;"`
	Symtoms      *string                     `gorm:"type:text"`
	Conclusion   *string                     `gorm:"type:text"`
	Prescription *string                     `gorm:"type:text"`
	IsRevisit    *bool                       `gorm:"type:bool;default:false"`
	IsPaid       *bool                       `gorm:"type:bool;default:false"`
	RevisitCode  *string                     `gorm:"type:varchar(64)"`
}

type Search struct {
	utils.DefaultSearchModel
}
