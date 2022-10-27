package api

import (
	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/authservice/src/internal/db"
	"github.com/minh1611/clinics_chain_management/authservice/src/internal/model"
	"github.com/minh1611/clinics_chain_management/authservice/src/pb/authpb"
)

var ApiServerSet = wire.NewSet(wire.Struct(new(ApiServer), "*"))

var (
	_ authpb.AuthServiceServer = (*ApiServer)(nil)
)

type ApiServer struct {
	authpb.UnimplementedAuthServiceServer `wire:"-"`
	Model                                 model.Server
	Repo                                  db.ServerRepo
}
