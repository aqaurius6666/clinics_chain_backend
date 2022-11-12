package model

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/staffservice/src/internal/db"
	"github.com/minh1611/clinics_chain_management/staffservice/src/services/authservice"
)

type Server interface {
	DoctorModel
	FeedbackModel
	GuestModel
	ManagerModel
	NurseModel
	PatientModel
	ProfileModel
	ReservationModel
	RevisitationModel
	ServiceTypeModel
	StaffWorkingTimeModel
	TreatmentRecordModel
}

type ServerModel struct {
	Ctx  context.Context
	Repo db.ServerRepo
	Auth authservice.Service
}

var ServerModelSet = wire.NewSet(wire.Bind(new(Server), new(*ServerModel)), wire.Struct(new(ServerModel), "*"))
