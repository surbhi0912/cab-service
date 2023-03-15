package entities

type Vehicle struct {
	CarNumber string
	CarType string
	IsAvailable bool
	Lat float64
	Long float64
	DriverId string // not tightly coupled, this logs which driver is currently driving the vehicle
}