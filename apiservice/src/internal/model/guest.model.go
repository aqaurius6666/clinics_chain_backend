package model

var (
	_ GuestModel = (*ServerModel)(nil)
)

type GuestModel interface {
}
