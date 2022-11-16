package api

import (
	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/authservice/src/internal/db"
	"github.com/minh1611/clinics_chain_management/authservice/src/internal/model"
	"github.com/minh1611/clinics_chain_management/authservice/src/pb/authpb"
)

var AuthServerSet = wire.NewSet(wire.Struct(new(AuthServer), "*"))

var (
	_ authpb.AuthServiceServer = (*AuthServer)(nil)
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer `wire:"-"`
	Model                                 model.Server
	Repo                                  db.ServerRepo
}
