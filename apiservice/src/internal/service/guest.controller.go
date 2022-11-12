package service

import "github.com/google/wire"

var GuestSet = wire.NewSet(wire.Struct(new(GuestController), "*"), wire.Struct(new(GuestService), "*"))

type GuestController struct {
	S *GuestService
}