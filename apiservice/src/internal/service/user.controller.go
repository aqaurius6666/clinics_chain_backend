package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/minh1611/go_structure/apiservice/src/internal/lib"
	"github.com/minh1611/go_structure/apiservice/src/pb"
)

var UserSet = wire.NewSet(wire.Struct(new(UserController), "*"), wire.Struct(new(UserService), "*"))

type UserController struct {
	S *UserService
}

type User struct {
	Name string
	Age  int16
	PhoneNumber string
	Email string
	Password string
}

func (s *UserController) HandleGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (s *UserController) HandlePost(ctx *gin.Context) {
	req := pb.User{}
	if err := lib.GetBody(ctx, &req); err != nil {
		lib.BadRequest(ctx, err)
		return
	}
	lib.Success(ctx, &req)
}
