package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/service"
)

var ApiServiceSet = wire.NewSet(wire.Struct(new(ApiService), "*"),
	gin.New,
	service.FeedbackSet,
	service.GuestSet,
	service.PatientSet,
	service.ProfileSet,
	service.ReservationSet,
	service.TreatmentRecordSet,
)

type ApiService struct {
	G               *gin.Engine
	Feedback        service.FeedbackController
	Guest           service.GuestController
	Patient         service.PatientController
	Profile         service.ProfileController
	Reservation     service.ReservationController
	TreatmentRecord service.TreatmentRecordController
}

func (s *ApiService) RegisterEndPoint() {
	// Patient
}
