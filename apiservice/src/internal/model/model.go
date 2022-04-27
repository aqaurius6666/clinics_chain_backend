package model

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/go_structure/apiservice/src/internal/db"
)

type Server interface {
	UserModel
}

type ServerModel struct {
	Ctx  context.Context
	Repo db.ServerRepo
}

var ServerModelSet = wire.NewSet(wire.Bind(new(Server), new(*ServerModel)), wire.Struct(new(ServerModel), "*"))
