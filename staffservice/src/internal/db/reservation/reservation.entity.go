package reservation

import (
	"time"

	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/db/reservation_comment"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/utils"
)

type Reservation struct {
	utils.DbBaseModel
	PatientID           uuid.UUID `gorm:"type:uuid"`
	DoctorID            uuid.UUID `gorm:"type:uuid"`
	IsConfirmed         *bool     `gorm:"type:bool;default:false"`
	Type                *string   `gorm:"type:varchar(64)"`
	Description         *string   `gorm:"type:text"`
	MeetingTime         time.Time
	ConfirmedBy         uuid.UUID `gorm:"type:uuid"`
	ReservationComments []*reservation_comment.ReservationComment
}

type Search struct {
	utils.DefaultSearchModel
}
