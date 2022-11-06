package feedback

import (
	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/treatment_record"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/utils"
)

type Feedback struct {
	utils.DbBaseModel
	TreatmentRecordID uuid.UUID `gorm:"type:uuid"`
	TreatmentRecord   *treatment_record.TreatmentRecord
	DoctorID          uuid.UUID `gorm:"type:uuid"`
	NurseID           uuid.UUID `gorm:"type:uuid"`
	Comment           *string   `gorm:"type:text"`
	Star              *int8     `gorm:"type:int4"`
	ImageUrl          *string   `gorm:"type:text"`
}

type Search struct {
	utils.DefaultSearchModel
}
