package db

import (
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/psql"
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
	NowFunc: func() time.Time {
		ti, _ := time.LoadLocation("Asia/Bangkok")
		return time.Now().In(ti)
	},
}

func ConnectGorm(dsn DBDsn) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(string(dsn)), gormConfig)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	fmt.Println("CONNECT DB SUCCESSFULLY")
	return db, err
}
