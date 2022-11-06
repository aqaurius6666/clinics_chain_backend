package reservation_comment

import (
	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/utils"
)

type ReservationComment struct {
	utils.DbBaseModel
	ReservationID uuid.UUID `gorm:"type:uuid"`
	CreatedBy     uuid.UUID `gorm:"type:uuid"`
	Content       *string   `gorm:"type:text"`
}

type Search struct {
	utils.DefaultSearchModel
}
