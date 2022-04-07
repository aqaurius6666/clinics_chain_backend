package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var PlayerSet = wire.NewSet(wire.Struct(new(PlayerController), "*"), wire.Struct(new(PlayerService), "*"))

type PlayerController struct {
	S PlayerService
}

func (s *PlayerController) HandleGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
