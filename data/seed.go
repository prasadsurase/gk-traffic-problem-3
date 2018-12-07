package data

import "github.com/prasadsurase/gk-traffic-problem-3/structs"

// CreateSeedData creates the seed places
func CreateSeedData() ([]*structs.Place, []*structs.Orbit, []*structs.Vehicle, []*structs.Weather) {
	silkRoad := &structs.Place{Name: "Silk Road"}
	hallitharam := &structs.Place{Name: "Hallitharam"}
	places := []*structs.Place{silkRoad, hallitharam}

	orbitOne := &structs.Orbit{Name: "Orbit 1", Source: silkRoad, Destination: hallitharam, Distance: 18, Craters: 20}
	orbitTwo := &structs.Orbit{Name: "Orbit 2", Source: silkRoad, Destination: hallitharam, Distance: 20, Craters: 10}
	orbits := []*structs.Orbit{orbitOne, orbitTwo}

	bike := &structs.Vehicle{Name: "Bike", Speed: 10, MinsPerCrater: 2}
	tuktuk := &structs.Vehicle{Name: "Tuktuk", Speed: 12, MinsPerCrater: 1}
	car := &structs.Vehicle{Name: "Car", Speed: 20, MinsPerCrater: 3}
	vehicles := []*structs.Vehicle{bike, tuktuk, car}

	sunny := &structs.Weather{Name: "Sunny", ChangeInCrater: 10, ChangeType: "Decrease", AllowedVehicles: []*structs.Vehicle{car, bike, tuktuk}}
	rainy := &structs.Weather{Name: "Rainy", ChangeInCrater: 20, ChangeType: "Increase", AllowedVehicles: []*structs.Vehicle{car, tuktuk}}
	windy := &structs.Weather{Name: "Windy", ChangeInCrater: 0, AllowedVehicles: []*structs.Vehicle{car, bike}}
	weathers := []*structs.Weather{sunny, rainy, windy}

	return places, orbits, vehicles, weathers
}
