package service

import "github.com/google/wire"

var PatientSet = wire.NewSet(wire.Struct(new(PatientController), "*"), wire.Struct(new(PatientService), "*"))

type PatientController struct {
	S *PatientService
}