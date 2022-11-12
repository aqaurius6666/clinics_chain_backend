package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	// "github.com/minh1611/clinics_chain_management/mailservice/src/internal/service"
)

var mailserviceSet = wire.NewSet(wire.Struct(new(mailservice), "*"),
	gin.New,
	//service.UserSet,
)

type mailservice struct {
	G *gin.Engine
	//User service.UserController
}

func (s *mailservice) RegisterEndPoint() {
	//api := s.G.Group("/api")
	//s.UserEndPoint(api)
}
