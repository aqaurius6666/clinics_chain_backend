package service

import "github.com/google/wire"

var NurseSet = wire.NewSet(wire.Struct(new(NurseController), "*"), wire.Struct(new(NurseService), "*"))

type NurseController struct {
	S *NurseService
}