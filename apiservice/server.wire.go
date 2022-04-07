// go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/go_example/apiservice/api"
)

type Server struct {
	ApiServer *api.ApiService
}

func InitMainServer(ctx context.Context) (*Server, error) {
	wire.Build(
		api.ApiServiceSet,
		wire.Struct(new(Server), "*"),
	)
	return &Server{}, nil
}
