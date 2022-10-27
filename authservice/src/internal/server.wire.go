//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/authservice/src/internal/api"
	"github.com/minh1611/clinics_chain_management/authservice/src/internal/db"
	"github.com/minh1611/clinics_chain_management/authservice/src/internal/model"
)

type Server struct {
	ApiServer *api.ApiServer
	MainRepo  db.ServerRepo
}

type ServerOptions struct {
	DBDsn db.DBDsn
}

func InitMainServer(ctx context.Context, opts ServerOptions) (*Server, error) {
	wire.Build(
		wire.FieldsOf(&opts, "DBDsn"),
		api.ApiServerSet,
		model.ServerModelSet,
		db.ServerRepoSet,
		wire.Struct(new(Server), "*"),
	)
	return &Server{}, nil
}
