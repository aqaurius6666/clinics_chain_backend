//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/go_structure/apiservice/src/internal/api"
	"github.com/minh1611/go_structure/apiservice/src/internal/db"
	"github.com/minh1611/go_structure/apiservice/src/internal/model"
)

type Server struct {
	ApiServer *api.ApiService
	MainRepo  db.ServerRepo
}

type ServerOptions struct {
	DBDsn db.DBDsn
}

func InitMainServer(ctx context.Context, opts ServerOptions) (*Server, error) {
	wire.Build(
		wire.FieldsOf(&opts, "DBDsn"),
		api.ApiServiceSet,
		model.ServerModelSet,
		db.ServerRepoSet,
		wire.Struct(new(Server), "*"),
	)
	return &Server{}, nil
}
