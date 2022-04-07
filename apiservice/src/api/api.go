package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/minh1611/go_structure/apiservice/src/service"
)

var ApiServiceSet = wire.NewSet(wire.Struct(new(ApiService), "*"),
	gin.New,
	service.PlayerSet,
)

type ApiService struct {
	G      *gin.Engine
	Player service.PlayerController
}

func (s *ApiService) RegisterEndPoint() {
	api := s.G.Group("/api")
	s.PlayerEndPoint(api)
}
