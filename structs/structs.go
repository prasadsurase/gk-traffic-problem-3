package structs

import (
	"fmt"
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

// DisplayDetails display the details
type DisplayDetails interface {
	DisplayDetails()
}

func (place *Place) DisplayDetails() {
	fmt.Println("Name:", place.Name)
}

func (orbit *Orbit) DisplayDetails() {
	fmt.Println("Name:", orbit.Name, ", Source:", orbit.Source.Name, ", Destination:", orbit.Destination.Name)
}

func (vehicle *Vehicle) DisplayDetails() {
	fmt.Println("Name:", vehicle.Name, ", Speed:", vehicle.Speed, ", Minutes per crater:", vehicle.MinsPerCrater)
}

func (weather *Weather) DisplayDetails() {
	fmt.Println("Name:", weather.Name)
}
