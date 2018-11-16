package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Place something goes here
type Place struct {
	Name string
}

// Orbit the distance between two places.
type Orbit struct {
	Name         string
	Source       *Place
	Destination  *Place
	Distance     int
	Craters      int
	CurrentSpeed int
}

// Vehicle details
type Vehicle struct {
	Name          string
	Speed         int
	MinsPerCrater int
}

//Weather details
type Weather struct {
	Name            string
	ChangeInCrater  int
	ChangeType      string
	AllowedVehicles []*Vehicle
}

// OrbitVehicle details
type OrbitVehicle struct {
	Orbit         *Orbit
	Vehicle       *Vehicle
	EstimatedTime float64
}

func main() {
	// fmt.Println("hola amigo")
	silkRoad := &Place{Name: "Silk Road"}
	hallitharam := &Place{Name: "Hallitharam"}
	places := []*Place{silkRoad, hallitharam}
	fmt.Println("Places: ", places)
	orbitOne := &Orbit{Name: "Orbit 1", Source: silkRoad, Destination: hallitharam, Distance: 18, Craters: 20}
	orbitTwo := &Orbit{Name: "Orbit 2", Source: silkRoad, Destination: hallitharam, Distance: 20, Craters: 10}
	orbits := []*Orbit{orbitOne, orbitTwo}
	fmt.Println("Orbits:", orbits)
	bike := &Vehicle{Name: "Bike", Speed: 10, MinsPerCrater: 2}
	tuktuk := &Vehicle{Name: "Tuktuk", Speed: 12, MinsPerCrater: 1}
	car := &Vehicle{Name: "Car", Speed: 20, MinsPerCrater: 3}
	vehicles := []*Vehicle{bike, tuktuk, car}
	fmt.Println("Vehicles: ", vehicles)

	sunny := &Weather{Name: "Sunny", ChangeInCrater: 10, ChangeType: "Decrease", AllowedVehicles: []*Vehicle{car, bike, tuktuk}}
	rainy := &Weather{Name: "Rainy", ChangeInCrater: 20, ChangeType: "Increase", AllowedVehicles: []*Vehicle{car, tuktuk}}
	windy := &Weather{Name: "Windy", ChangeInCrater: 0, AllowedVehicles: []*Vehicle{car, bike}}

	weathers := []*Weather{sunny, rainy, windy}
	fmt.Println("Weathers: ", weathers)
	var orbitsSpeeds []*OrbitVehicle

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Specify the Weather (Sunny, Rainy, Windy:")
	inputWeather, _ := reader.ReadString([]byte("\n")[0])
	inputWeather = strings.TrimSuffix(inputWeather, "\n")

	for _, orbit := range orbits {
		fmt.Println("Enter current max speed for", orbit.Name)
		fmt.Scan(&orbit.CurrentSpeed)
	}

	switch inputWeather {
	case "Sunny":
		for _, vehicle := range sunny.AllowedVehicles {
			// fmt.Println("Vehicle:", vehicle)
			for _, orb := range orbits {
				estimatedTime := calculateEstimatedTime(sunny, orb, vehicle)
				ov := &OrbitVehicle{Vehicle: vehicle, Orbit: orb, EstimatedTime: estimatedTime}
				fmt.Println("OrbitVehicle: Name:", orb.Name, ", Vehicle:", vehicle.Name, ", Estimated Time:", estimatedTime)
				orbitsSpeeds = append(orbitsSpeeds, ov)
			}
		}
	case "Rainy":
		for _, v := range rainy.AllowedVehicles {
			fmt.Println("Vehicle:", v)
		}
	case "Windy":
		for _, v := range windy.AllowedVehicles {
			fmt.Println("Vehicle:", v)
		}
	}
}

func calculateEstimatedTime(weather *Weather, orbit *Orbit, vehicle *Vehicle) float64 {
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
