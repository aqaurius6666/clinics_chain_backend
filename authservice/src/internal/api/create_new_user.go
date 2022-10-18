package api

import (
	"context"
	"fmt"

	"github.com/minh1611/go_structure/authservice/src/pb/authpb"
)

func (s *ApiServer) CreateNewUser(ctx context.Context, user *authpb.NewUser) (*authpb.User, error) {
	
	return &authpb.User{
		Name: user.Name,
		Id: 1,
	}, nil
}
