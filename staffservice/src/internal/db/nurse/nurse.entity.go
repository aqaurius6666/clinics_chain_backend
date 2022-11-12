package nurse

import (
	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/db/profile"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/db/staff_working_time"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/db/treatment_record"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/utils"
)

type Nurse struct {
	utils.DbBaseModel
	BranchID         uuid.UUID        `gorm:"type:uuid"`
	ProfileID        uuid.UUID        `gorm:"type:uuid"`
	Profile          *profile.Profile `gorm:"foreginKey:ProfileID"`
	FirstName        *string          `gorm:"type:varchar(64)"`
	LastName         *string          `gorm:"type:varchar(64)"`
	Age              *int             `gorm:"type:int4;default:0"`
	PhoneNumber      *string          `gorm:"type:varchar(64)"`
	Email            *string          `gorm:"type:varchar(64)"`
	Address          *string          `gorm:"type:varchar(64)"`
	Role             *string          `gorm:"type:varchar(16)"`
	TreatmentRecords []*treatment_record.TreatmentRecord
	Timetables       []*staff_working_time.StaffWorkingTime
}

type Search struct {
	utils.DefaultSearchModel
}
