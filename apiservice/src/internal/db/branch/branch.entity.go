package branch

import (
	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/manager"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/utils"
)

type Branch struct {
	utils.DbBaseModel
	BranchManagerID uuid.UUID        `gorm:"type:uuid"`
	Address         *string          `gorm:"type:text"`
	Manager         *manager.Manager `gorm:"foreignKey:BranchManagerID"`
}

type Search struct {
	utils.DefaultSearchModel
}