package doctor

import (
	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/db/reservation"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/db/staff_working_time"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/db/treatment_record"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/utils"
)

type Doctor struct {
	utils.DbBaseModel
	BranchID         uuid.UUID `gorm:"type:uuid"`
	FirstName        *string   `gorm:"type:varchar(64)"`
	LastName         *string   `gorm:"type:varchar(64)"`
	Age              *int      `gorm:"type:int4;default:0"`
	PhoneNumber      *string   `gorm:"type:varchar(64)"`
	Email            *string   `gorm:"type:varchar(64)"`
	Address          *string   `gorm:"type:varchar(64)"`
	Role             *string   `gorm:"type:varchar(16)"`
	TreatmentRecords []*treatment_record.TreatmentRecord
	Timetables       []*staff_working_time.StaffWorkingTime
	Reservations     []*reservation.Reservation
}

type Search struct {
	utils.DefaultSearchModel
}
