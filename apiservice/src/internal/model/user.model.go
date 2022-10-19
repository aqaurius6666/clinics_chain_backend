package model

import (
	"context"

	"github.com/minh1611/go_structure/apiservice/src/pb/authpb"
)

var (
	_ UserModel = (*ServerModel)(nil)
)

type UserModel interface {
	CreateNewUser(ctx context.Context, name string, age int) (*authpb.User, error)
}

func (s *ServerModel) CreateNewUser(ctx context.Context, name string, age int) (*authpb.User, error) {
	user, err := s.Auth.CreateNewUser(ctx, name, age)
	if err != nil {
		return nil, err
	}
	return user, nil
}
