package manager

import "github.com/minh1611/clinics_chain_management/mailservice/src/internal/utils"

type Manager struct {
	utils.DbBaseModel
}

type Search struct {
	utils.DefaultSearchModel
}
