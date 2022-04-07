package api

import (
	"github.com/gin-gonic/gin"
)

func (s *ApiService) PlayerEndPoint(g *gin.RouterGroup) {
	playerGroup := g.Group("/player")

	playerGroup.GET("/", s.Player.HandleGet)
}
