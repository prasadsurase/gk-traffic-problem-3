package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/prasadsurase/gk-traffic-problem-3/calculations"

	"github.com/prasadsurase/gk-traffic-problem-3/data"
	"github.com/prasadsurase/gk-traffic-problem-3/search"
	"github.com/prasadsurase/gk-traffic-problem-3/structs"
)

func main() {
	places, orbits, vehicles, weathers := data.CreateSeedData()

	var option int
	// reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please select an option.")
		fmt.Println("1) Display data.")
		fmt.Println("2) Go directly from source to destination.")
		fmt.Println("3) Go from source to destination via other places.")
		fmt.Println("4) Exit")
		fmt.Scan(&option)
		switch option {
		case 1:
			fmt.Println("Places: ", places)
			fmt.Println("Orbits:", orbits)
			fmt.Println("Vehicles: ", vehicles)
			fmt.Println("Weathers: ", weathers)
		case 2:
			handleFirstCase(places, orbits, vehicles, weathers)
		case 3:
			// handleSecondCase()
		case 4:
			return
		default:
			fmt.Println("Wrong choice. Try Again.")
		}
	}
}

func handleFirstCase(places []*structs.Place, orbits []*structs.Orbit, vehicles []*structs.Vehicle, weathers []*structs.Weather) {
	reader := bufio.NewReader(os.Stdin)
	// accept the source. TODO: refactor to handle unexpected places.
	fmt.Println("Enter Source:")
	sourceLocation, _ := reader.ReadString([]byte("\n")[0])
	sourceLocation = strings.TrimSuffix(sourceLocation, "\n")
	source, err := search.Place(sourceLocation, places)
	if err != nil {

	}
	// accept the destination. TODO: refactor to handle unexpected places.
	fmt.Println("Enter Destination:")
	destinationLocation, _ := reader.ReadString([]byte("\n")[0])
	destinationLocation = strings.TrimSuffix(destinationLocation, "\n")
	destination, err := search.Place(destinationLocation, places)

	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)

	fmt.Println("Specify the Weather (Sunny, Rainy, Windy:")
	inputWeather, _ := reader.ReadString([]byte("\n")[0])
	inputWeather = strings.TrimSuffix(inputWeather, "\n")

	var selectedWeather *structs.Weather
	// accept the weather.
	for _, w := range weathers {
		if w.Name == inputWeather {
			selectedWeather = w
			break
		}
	}
	fmt.Println("Selected Weather:", selectedWeather)

	// find direct orbits from source to destination
	directOrbits := getDirectOrbits(orbits, source, destination)
	fmt.Println("Direct Orbits:", directOrbits)
	for _, orbit := range directOrbits {
		fmt.Println("Enter current max speed for", orbit.Name)
		fmt.Scan(&orbit.CurrentSpeed)
	}
	var orbitsSpeeds []*structs.OrbitVehicle

	// accept current orbit speed for direct orbits
	for _, vehicle := range selectedWeather.AllowedVehicles {
		for _, orb := range directOrbits {
			estimatedTime := calculations.TravelTime(selectedWeather, orb, vehicle)
			ov := &structs.OrbitVehicle{Vehicle: vehicle, Orbit: orb, EstimatedTime: estimatedTime}
			fmt.Println("OrbitVehicle: Name:", orb.Name, ", Vehicle:", vehicle.Name, ", Estimated Time:", estimatedTime, " minutes")
			orbitsSpeeds = append(orbitsSpeeds, ov)
		}
	}

	//sort the orbits as per the calculated time to reach destination
	sort.Slice(orbitsSpeeds[:], func(i, j int) bool {
		return orbitsSpeeds[i].EstimatedTime < orbitsSpeeds[j].EstimatedTime
	})
	fmt.Println("Best orbit details:")
	orbSpeed := orbitsSpeeds[0]
	fmt.Println("OrbitVehicle: Name:", orbSpeed.Orbit.Name, ", Vehicle:", orbSpeed.Vehicle.Name, ", Estimated Time:", orbSpeed.EstimatedTime, " minutes")
}

// a func where the source and destination have no intermediary places.
func getDirectOrbits(orbits []*structs.Orbit, source *structs.Place, destination *structs.Place) []*structs.Orbit {
	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
	filteredOrbits := []*structs.Orbit{}
	for _, orbit := range orbits {
		if ((orbit.Source == source) && (orbit.Destination == destination)) || ((orbit.Source == destination) && (orbit.Destination == source)) {
			filteredOrbits = append(filteredOrbits, orbit)
		}
	}
	return filteredOrbits
}
