package db

import "github.com/minh1611/go_structure/authservice/src/internal/db/user"

type DBDsn string

type ServerRepo interface {
	user.UserRepo
}
