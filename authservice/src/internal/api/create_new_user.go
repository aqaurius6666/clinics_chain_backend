package api

import (
	"context"
	"fmt"

	"github.com/minh1611/go_structure/authservice/src/pb"
)

func (s *ApiServer) CreateNewUser(ctx context.Context, user *pb.NewUser) (*pb.User, error) {
	fmt.Println(user.Name)
	return &pb.User{
		Name: user.Name,
	}, nil
}
