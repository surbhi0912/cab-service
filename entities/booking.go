package entities

type Booking struct {
	BookingId string
	RiderUserId string
	CarNumber string
	StartTime int64
	EndTime int64
	Status string
}