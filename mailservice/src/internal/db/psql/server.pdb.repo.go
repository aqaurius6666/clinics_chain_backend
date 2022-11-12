package psql

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/branch"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/doctor"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/feedback"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/guest"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/manager"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/nurse"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/patient"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/profile"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/profile_timeline"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/reservation"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/reservation_comment"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/service_type"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/staff_working_time"
	"github.com/minh1611/clinics_chain_management/mailservice/src/internal/db/treatment_record"
	"gorm.io/gorm"
)

var (
	timeout = 2 * time.Second
)

var CDBRepoSet = wire.NewSet(wire.Struct(new(ServerCDBRepo), "*"), InterfaceProvider)

type DBInterfaces []interface{}

func InterfaceProvider() DBInterfaces {
	return DBInterfaces{
		branch.Branch{},
		doctor.Doctor{},
		feedback.Feedback{},
		guest.Guest{},
		manager.Manager{},
		nurse.Nurse{},
		patient.Patient{},
		profile.Profile{},
		profile_timeline.ProfileTimeline{},
		reservation.Reservation{},
		reservation_comment.ReservationComment{},
		service_type.ServiceType{},
		staff_working_time.StaffWorkingTime{},
		treatment_record.TreatmentRecord{},
	}
}

type ServerCDBRepo struct {
	Db         *gorm.DB
	Context    context.Context
	Interfaces DBInterfaces
	//Logger     *logrus.Logger
}
