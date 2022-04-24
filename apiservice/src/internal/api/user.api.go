package api

import (
	"github.com/gin-gonic/gin"
)

func (s *ApiService) UserEndPoint(g *gin.RouterGroup) {
	userGroup := g.Group("/user")

	userGroup.GET("/", s.User.HandleGet)
}
