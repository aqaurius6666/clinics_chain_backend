package model

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/go_structure/apiservice/src/internal/db"
	"github.com/minh1611/go_structure/apiservice/src/services/authservice"
)

type Server interface {
	UserModel
}

type ServerModel struct {
	Ctx  context.Context
	Repo db.ServerRepo
	Auth authservice.Service
}

var ServerModelSet = wire.NewSet(wire.Bind(new(Server), new(*ServerModel)), wire.Struct(new(ServerModel), "*"))
