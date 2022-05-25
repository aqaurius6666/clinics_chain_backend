package service

import (
	"context"

	"github.com/minh1611/go_structure/apiservice/src/internal/model"
)

type UserService struct {
	Model model.Server
}

func (s *UserService) AddUser(ctx context.Context, req User) {
	
}
