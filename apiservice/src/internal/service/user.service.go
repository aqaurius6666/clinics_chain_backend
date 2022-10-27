package service

import (
	"context"

	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/model"
	"github.com/minh1611/clinics_chain_management/apiservice/src/pb"
)

type UserService struct {
	Model model.Server
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.Model.CreateNewUser(ctx, &req.Name, nil)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Name: *user.Name,
		Age:  99,
		//Id:   user.ID.String(),
	}, nil
}
