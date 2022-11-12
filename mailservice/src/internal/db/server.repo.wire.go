package db

import (
	"fmt"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/psql"
	"golang.org/x/xerrors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ServerRepoSet = wire.NewSet(ConnectGorm, wire.Bind(new(ServerRepo), new(*psql.ServerCDBRepo)), psql.CDBRepoSet)

var gormConfig = &gorm.Config{
	//SkipDefaultTransaction:                   true,
	DisableAutomaticPing:                     true,
	PrepareStmt:                              true,
	DisableForeignKeyConstraintWhenMigrating: true,
}

func ConnectGorm(dsn DBDsn) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(string(dsn)), gormConfig)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	fmt.Println("CONNECT DB SUCCESSFULLY")
	return db, err
}
