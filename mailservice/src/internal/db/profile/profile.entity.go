package profile

import (
	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/doctor"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/utils"
)

type Profile struct {
	utils.DbBaseModel
	DoctorID      uuid.UUID `gorm:"type:uuid"`
	Doctor        *doctor.Doctor
	AcademicLevel *string `gorm:"type:text"`
	Introduction  *string `gorm:"type:text"`
}

type Search struct {
	utils.DefaultSearchModel
}
