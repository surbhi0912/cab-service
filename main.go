package main

import (
	"cab-service/controllers"
	"cab-service/entities"
	"fmt"
	"time"
)


func main(){
	// driverData := controllers.SaveDriver()
	// fmt.Println(driverData)

	var rider entities.Rider
	rider = entities.Rider{
		RiderId : "1",
		Name :"Surbhi",
		CountryCode : "+91",
		PhoneNumber : "7077100596",
	}
	controllers.SaveRider(rider)
	rider = entities.Rider{
		RiderId : "2",
		Name :"Nidhi",
		CountryCode : "+91",
		PhoneNumber : "1234567890",
	}
	controllers.SaveRider(rider)
	// fmt.Println(controllers.RiderMap)

	var driver entities.Driver 
	driver = entities.Driver{
		DriverId : "7",
		Name :"ABC",
		CountryCode : "+91",
		PhoneNumber : "2342342345",
	}
	controllers.SaveDriver(driver)
	driver = entities.Driver{
		DriverId : "9",
		Name :"DEF",
		CountryCode : "+91",
		PhoneNumber : "5675675678",
	}
	controllers.SaveDriver(driver)

	var vehicle entities.Vehicle
	vehicle = entities.Vehicle{
		CarNumber : "XY0001",
		CarType : "SUV",
		IsAvailable : true,
		Lat : 76.1234567,
		Long : 45.567799,
		DriverId : "7",
	}
	controllers.SaveVehicle(vehicle)
	vehicle = entities.Vehicle{
		CarNumber : "CZ0005",
		CarType : "Sedan",
		IsAvailable : true,
		Lat : 76.1234567,
		Long : 45.567799,
		DriverId : "9",
	}
	controllers.SaveVehicle(vehicle)
	// fmt.Println("After saving", controllers.VehicleMap)

	controllers.UpdateLocation(&vehicle, 77.12345, 44.12345)
	// fmt.Println("main", vehicle)
	// fmt.Println("After updating location", controllers.VehicleMap)

	// controllers.FindVehicle(77.12346, 44.12346, 15)
	// fmt.Println("Found vehicles", controllers.AvailableVehicles)

	bookingId, status := controllers.MakeBooking("1", 77.12346, 44.12346)
	fmt.Println(bookingId, status, controllers.RiderMap, controllers.VehicleMap)

	bookingId, status = controllers.MakeBooking("2", 77.12347, 44.12347)
	fmt.Println(bookingId, status, controllers.RiderMap, controllers.VehicleMap)

	controllers.EndTrip(bookingId, time.Now().UnixNano())
	fmt.Println(controllers.BookingMap, controllers.VehicleMap)

}