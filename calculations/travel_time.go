package calculations

import "github.com/prasadsurase/gk-traffic-problem-3/structs"

// TravelTime returns the estimated time for all provided orbits for all provided vehiles
// between source and destination
func TravelTime(weather *structs.Weather, orbit *structs.Orbit, vehicle *structs.Vehicle) float64 {
	maxSpeed := 0
	if orbit.CurrentSpeed <= vehicle.Speed {
		maxSpeed = orbit.CurrentSpeed
	} else {
		maxSpeed = vehicle.Speed
	}

	craters := float64(orbit.Craters)
	switch weather.ChangeType {
	case "Decrease":
		craters -= (craters / 100) * float64(weather.ChangeInCrater)
	case "Increase":
		craters += (craters / 100) * float64(weather.ChangeInCrater)
	}

	timeForCraters := craters * float64(vehicle.MinsPerCrater)
	return float64((orbit.Distance*60)/maxSpeed) + timeForCraters
}
