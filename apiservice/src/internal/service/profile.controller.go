package service

import "github.com/google/wire"

var ProfileSet = wire.NewSet(wire.Struct(new(ProfileController), "*"), wire.Struct(new(ProfileService), "*"))

type ProfileController struct {
	S *ProfileService
}
