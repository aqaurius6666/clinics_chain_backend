package profile_timeline

import (
	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/utils"
)

type ProfileTimeline struct {
	utils.DbBaseModel
	ProfileID   uuid.UUID
	TimeFrom    *string `gorm:"type:varchar(256)"`
	TimeTo      *string `gorm:"type:varchar(256)"`
	Description *string `gorm:"type:text"`
}

type Search struct {
	utils.DefaultSearchModel
}
