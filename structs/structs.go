package structs

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
