package api

import (
	"github.com/google/wire"
	"github.com/minh1611/go_structure/authservice/src/internal/db"
	"github.com/minh1611/go_structure/authservice/src/internal/model"
	"github.com/minh1611/go_structure/authservice/src/pb"
)

var ApiServerSet = wire.NewSet(wire.Struct(new(ApiServer), "*"))

var (
	_ pb.AuthServiceServer = (*ApiServer)(nil)
)

type ApiServer struct {
	pb.UnimplementedAuthServiceServer `wire:"-"`
	Model                             model.Server
	Repo                              db.ServerRepo
}
