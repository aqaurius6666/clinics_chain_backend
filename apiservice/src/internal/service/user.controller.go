package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/lib"
	"github.com/minh1611/clinics_chain_management/apiservice/src/pb"
)

var UserSet = wire.NewSet(wire.Struct(new(UserController), "*"), wire.Struct(new(UserService), "*"))

type UserController struct {
	S *UserService
}

type User struct {
	Name        string
	Age         int16
	PhoneNumber string
	Email       string
	Password    string
}

func (s *UserController) HandleGet(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (s *UserController) HandlePost(g *gin.Context) {
	req := pb.User{}
	if err := lib.GetBody(g, &req); err != nil {
		lib.BadRequest(g, err)
		return
	}
	res, err := s.S.CreateUser(g, &req)
	if err != nil {
		lib.BadRequest(g, err)
	}
	lib.Success(g, &res)
}
