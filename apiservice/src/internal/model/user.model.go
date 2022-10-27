package model

import (
	"context"

	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/user"
	//"github.com/minh1611/clinics_chain_management/apiservice/src/pb/authpb"
)

var (
	_ UserModel = (*ServerModel)(nil)
)

type UserModel interface {
	CreateNewUser(ctx context.Context, name *string, age *int) (*user.User, error)
}

func (s *ServerModel) CreateNewUser(ctx context.Context, name *string, age *int) (*user.User, error) {
	user, err := s.Repo.InsertUser(ctx, &user.User{
		Name: name,
		Age:  age,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
