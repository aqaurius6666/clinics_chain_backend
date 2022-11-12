package service

import "github.com/google/wire"

var FeedbackSet = wire.NewSet(wire.Struct(new(FeedbackController), "*"), wire.Struct(new(FeedbackService), "*"))

type FeedbackController struct {
	S *FeedbackService
}
