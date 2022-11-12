package staff_working_time

import (
	"time"

	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/utils"
)

type StaffWorkingTime struct {
	utils.DbBaseModel
	NurseID  uuid.UUID `gorm:"type:uuid"`
	DoctorID uuid.UUID `gorm:"type:uuid"`
	Date     *string   `gorm:"type:varchar(16)"`
	TimeFrom time.Time
	TimeTo   time.Time
}

type Search struct {
	utils.DefaultSearchModel
}
