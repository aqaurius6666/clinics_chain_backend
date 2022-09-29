//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	//"github.com/minh1611/go_structure/authservice/src/internal/api"
	//"github.com/minh1611/go_structure/authservice/src/internal/db"
	//"github.com/minh1611/go_structure/authservice/src/internal/model"
)

type Server struct {
	ApiServer *api.AuthService
	MainRepo  db.ServerRepo
}

type ServerOptions struct {
	DBDsn db.DBDsn
}

func InitMainServer(ctx context.Context, opts ServerOptions) (*Server, error) {
	wire.Build(
		wire.FieldsOf(&opts, "DBDsn"),
		api.AuthServiceSet,
		model.ServerModelSet,
		db.ServerRepoSet,
		wire.Struct(new(Server), "*"),
	)
	return &Server{}, nil
}
