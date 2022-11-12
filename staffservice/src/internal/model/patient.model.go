package model

var (
	_ PatientModel = (*ServerModel)(nil)
)

type PatientModel interface {
}
