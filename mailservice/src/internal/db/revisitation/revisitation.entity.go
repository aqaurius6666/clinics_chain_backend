package revisitation

import (
	"time"

	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/utils"
)

type Revisitation struct {
	utils.DbBaseModel
	PreviousTreatmentRecordID uuid.UUID `gorm:"type:uuid"`
	Code                      *string   `gorm:"type:varchar(64)"`
	Time                      time.Time
}

type Search struct {
	utils.DefaultSearchModel
}
