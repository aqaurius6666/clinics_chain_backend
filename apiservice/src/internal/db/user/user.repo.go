package user

import "context"

type UserRepo interface {
	InsertUser(context.Context, *User) (*User, error)
}
