package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/minh1611/go_structure/authservice/src/internal/service"
)

var AuthServiceSet = wire.NewSet(wire.Struct(new(AuthService), "*"),
	gin.New,
	service.UserSet,
)

type AuthService struct {
	G    *gin.Engine
	User service.UserController
}

func (s *AuthService) RegisterEndPoint() {
	api := s.G.Group("/api")
	s.UserEndPoint(api)
}
