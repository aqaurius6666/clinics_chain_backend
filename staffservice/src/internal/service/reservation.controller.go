package service

import "github.com/google/wire"

var ReservationSet = wire.NewSet(wire.Struct(new(ReservationController), "*"), wire.Struct(new(ReservationService), "*"))

type ReservationController struct {
	S *ReservationService
}