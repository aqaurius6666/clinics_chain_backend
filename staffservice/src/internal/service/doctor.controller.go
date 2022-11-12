package service

import "github.com/google/wire"

var DoctorSet = wire.NewSet(wire.Struct(new(DoctorController), "*"), wire.Struct(new(DoctorService), "*"))

type DoctorController struct {
	S *DoctorService
}