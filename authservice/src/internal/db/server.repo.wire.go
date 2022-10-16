package db

import (
	"fmt"

	"github.com/google/wire"
	"github.com/minh1611/go_structure/authservice/src/internal/db/my"
	"golang.org/x/xerrors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ServerRepoSet = wire.NewSet(ConnectGorm, wire.Bind(new(ServerRepo), new(*my.ServerCDBRepo)), my.CDBRepoSet)

var gormConfig = &gorm.Config{
	SkipDefaultTransaction:                   true,
	DisableAutomaticPing:                     true,
	PrepareStmt:                              true,
	DisableForeignKeyConstraintWhenMigrating: true,
}

func ConnectGorm(dsn DBDsn) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(string(dsn)), gormConfig)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	fmt.Println("CONNECT DB SUCCESSFULLY")
	return db, err
}
