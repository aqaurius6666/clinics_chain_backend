package service

import (
	"context"

	"github.com/minh1611/clinics_chain_management/authservice/src/internal/model"
	"github.com/minh1611/clinics_chain_management/authservice/src/pb/authpb"
)

type UserService struct {
	Model model.Server
}

func (s *UserService) AddUser(ctx context.Context, req *authpb.User) {

}
