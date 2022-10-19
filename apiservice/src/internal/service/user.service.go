package service

import (
	"context"

	"github.com/minh1611/go_structure/apiservice/src/internal/model"
	"github.com/minh1611/go_structure/apiservice/src/pb"
)

type UserService struct {
	Model model.Server
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.Model.CreateNewUser(ctx, req.Name, int(req.Age))
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Name: user.Name,
		Age:  user.Age,
		Id:   user.Id,
	}, nil
}
