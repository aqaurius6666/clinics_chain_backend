package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/service"
)

var StaffServiceSet = wire.NewSet(wire.Struct(new(StaffService), "*"),
	gin.New,
	service.DoctorSet,
	service.NurseSet,
	service.ProfileSet,
	service.ReservationSet,
	service.TreatmentRecordSet,
)

type StaffService struct {
	G *gin.Engine
	Doctor service.DoctorController
	Nurse service.NurseController
	Profile service.ProfileController
	Reservation service.ReservationController
	TreatmentReord service.TreatmentRecordController
}

func (s *StaffService) RegisterEndPoint() {
}
