package controllers

import (
	"cab-service/entities"
	// "fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	MaxRadius float64 = 15
)

func generateBookingId() string {
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Int())
	// fmt.Println(rand.Int())
	return strconv.FormatInt(int64(rand.Int()), 10)
}

func MakeBooking(userId string, userLat float64, userLong float64) (string, bool){
	vehicleId, found := FindVehicle(userLat, userLong, MaxRadius)
	if found { //nearest vehicle found
		// fmt.Println(vehicleId)
		bId := generateBookingId()
		if _, found := BookingMap[bId]; found{
			log.Fatal("Booking Id already exists")
			return "", false
		} else{
			//make this vehicleId IsAvailable as false after assignment done
			if x, exist := VehicleMap[vehicleId]; exist{
				x.IsAvailable = false
				VehicleMap[vehicleId] = x
				//make a booking
				var booking entities.Booking
				booking = entities.Booking{
					BookingId : bId,
					RiderUserId: userId,
					CarNumber : vehicleId,
					StartTime : time.Now().UnixNano(),
					Status : "Ongoing",
				}
				BookingMap[bId] = booking
				//Add booking id to Rider BookingHistory
				if r, e := RiderMap[userId]; e{
					r.BookingHistory = append(r.BookingHistory, bId)
					RiderMap[userId] = r
				}
				return bId, true
			}
			return "", false
		}
	} else {
		return "",false
	}
}

func EndTrip(bookingId string, endTimeStamp int64) {
	if v, found := BookingMap[bookingId]; found{
		v.EndTime = endTimeStamp
		BookingMap[bookingId] = v
		if x, e := VehicleMap[v.CarNumber]; e{
			x.IsAvailable = true
			VehicleMap[v.CarNumber] = x
		}
	}
}