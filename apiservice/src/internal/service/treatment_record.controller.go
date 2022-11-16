package service

import (
	"fmt"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/apiservice/src/pb"
)


var TreatmentRecordSet = wire.NewSet(wire.Struct(new(TreatmentRecordController), "*"), wire.Struct(new(TreatmentRecordService), "*"))

type TreatmentRecordController struct {
	S *TreatmentRecordService
}

func SOS() {
	//err := req.Vali
	a := new(pb.PostUserClient)
	fmt.Println(a)
}