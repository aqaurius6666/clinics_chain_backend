package psql

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/user"
	"gorm.io/gorm"
)

var (
	timeout = 2 * time.Second
)

var CDBRepoSet = wire.NewSet(wire.Struct(new(ServerCDBRepo), "*"), InterfaceProvider)

type DBInterfaces []interface{}

func InterfaceProvider() DBInterfaces {
	return DBInterfaces{
		user.User{},
	}
}

type ServerCDBRepo struct {
	Db         *gorm.DB
	Context    context.Context
	Interfaces DBInterfaces
	//Logger     *logrus.Logger
}
