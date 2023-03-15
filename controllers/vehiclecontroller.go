package controllers

import (
	"cab-service/entities"
	// "fmt"
	"math"
)

func calculateDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	a := (lat2 - lat1)
	b := (lng2 - lng1)
	return math.Sqrt(a*a + b*b)
}

func UpdateLocation(vehicle *entities.Vehicle, xNew float64, yNew float64){
	vehicle.Lat = xNew
	vehicle.Long = yNew
	VehicleMap[vehicle.CarNumber] = *vehicle
	// fmt.Println(VehicleMap[vehicle.CarNumber])
}

func FindVehicle(userLat float64, userLong float64, radius float64) (string, bool) {
	var distance float64
	minDistance := math.Inf(1)
	var nearestVehicleId string
	for key, value := range VehicleMap {
		if (value.IsAvailable == true){
			distance = calculateDistance(userLat, userLong, value.Lat, value.Long)
			// fmt.Println("distance", distance)
			if(distance <= radius){ // can be assigned to user
				if(distance < minDistance){
					minDistance = distance
					nearestVehicleId = key
				}
			}
		}
	}
	if (nearestVehicleId == ""){
		return "", false
	}
	return nearestVehicleId, true
}