package api

import (
	"context"

	"github.com/minh1611/clinics_chain_management/authservice/src/pb/authpb"
)

func (s *ApiServer) CreateNewUser(ctx context.Context, user *authpb.NewUser) (*authpb.User, error) {
	return &authpb.User{
		Name: user.Name,
		Age:  user.Age,
		Id:   100,
	}, nil
}
