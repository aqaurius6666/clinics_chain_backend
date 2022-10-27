//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/api"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/model"
	"github.com/minh1611/clinics_chain_management/apiservice/src/services/authservice"
)

type Server struct {
	ApiServer *api.ApiService
	MainRepo  db.ServerRepo
}

type ServerOptions struct {
	DBDsn           db.DBDsn
	AuthServiceAddr authservice.AuthServiceAddr
}

func InitMainServer(ctx context.Context, opts ServerOptions) (*Server, error) {
	wire.Build(
		wire.FieldsOf(&opts, "DBDsn", "AuthServiceAddr"),
		api.ApiServiceSet,
		model.ServerModelSet,
		db.ServerRepoSet,
		authservice.Set,
		wire.Struct(new(Server), "*"),
	)
	return &Server{}, nil
}
