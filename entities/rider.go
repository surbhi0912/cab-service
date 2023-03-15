package entities

type Rider struct {
	RiderId string
	Name string
	CountryCode string
	PhoneNumber string
	BookingHistory []string //slice
}