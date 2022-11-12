package service

import "github.com/google/wire"

var TreatmentRecordSet = wire.NewSet(wire.Struct(new(TreatmentRecordController), "*"), wire.Struct(new(TreatmentRecordService), "*"))

type TreatmentRecordController struct {
	S *TreatmentRecordService
}