package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/minh1611/go_structure/apiservice/src/internal/service"
)

var ApiServiceSet = wire.NewSet(wire.Struct(new(ApiService), "*"),
	gin.New,
	service.UserSet,
)

type ApiService struct {
	G    *gin.Engine
	User service.UserController
}

func (s *ApiService) RegisterEndPoint() {
	api := s.G.Group("/api")
	s.UserEndPoint(api)
}
