package model

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/authservice/src/internal/db"
)

type Server interface {
	UserModel
}

type ServerModel struct {
	Ctx  context.Context
	Repo db.ServerRepo
}

var ServerModelSet = wire.NewSet(wire.Bind(new(Server), new(*ServerModel)), wire.Struct(new(ServerModel), "*"))
