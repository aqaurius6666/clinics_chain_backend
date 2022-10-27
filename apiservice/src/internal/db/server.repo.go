package db

import "github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/user"

type DBDsn string

type ServerRepo interface {
	user.UserRepo
}
