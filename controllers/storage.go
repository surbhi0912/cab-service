package controllers

import (
	"cab-service/entities"
	// "fmt"
)


var DriverMap = make(map[string]entities.Driver)
var RiderMap = make(map[string]entities.Rider)
var VehicleMap = make(map[string]entities.Vehicle)
var BookingMap = make(map[string]entities.Booking)


func SaveDriver(driver entities.Driver) {
	DriverMap[driver.DriverId] = driver
	// fmt.Println(DriverMap)
}

func SaveRider(rider entities.Rider) {
	// fmt.Println("Hello saverider")
	// fmt.Println(rider)
	RiderMap[rider.RiderId] = rider
	// fmt.Println(RiderMap)
}

func SaveVehicle(vehicle entities.Vehicle) {
	VehicleMap[vehicle.CarNumber] = vehicle
	// fmt.Println(VehicleMap)
}