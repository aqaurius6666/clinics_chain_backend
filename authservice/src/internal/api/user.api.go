package api

import (
	"github.com/gin-gonic/gin"
)

func (s *AuthService) UserEndPoint(g *gin.RouterGroup) {
	userGroup := g.Group("/user")

	userGroup.GET("/", s.User.HandleGet)
	userGroup.POST("/", s.User.HandlePost)
}
