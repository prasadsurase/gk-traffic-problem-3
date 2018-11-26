package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
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

  // Accept source and destination
  fmt.Println("Specify the source location: ")
  for _, src := range places {
    fmt.Println(src.Name)
  }
  sourceLocation, _ := reader.ReadString([]byte("\n")[0])
  sourceLocation = strings.TrimSuffix(sourceLocation, "\n")
  source := getPlace(sourceLocation, places)
  // TODO: make sure that source is available in the list

  fmt.Println("Specify the destination location: ")
  for _, dest := range places {
    fmt.Println(dest.Name)
  }

  destinationLocation, _ := reader.ReadString([]byte("\n")[0])
  destinationLocation = strings.TrimSuffix(destinationLocation, "\n")
  destination := getPlace(destinationLocation, places)
  // TODO: make sure that destination is available in the list and not != source

  // filter out orbits that take from source to destination.
  firstCaseOrbits := handleFirstCase(orbits, source, destination)

  fmt.Println("Specify the Weather (Sunny, Rainy, Windy:")
  inputWeather, _ := reader.ReadString([]byte("\n")[0])
  inputWeather = strings.TrimSuffix(inputWeather, "\n")

  var selectedWeather *Weather

  for _, w := range weathers {
    if w.Name == inputWeather {
      selectedWeather = w
      break
    }
  }

  fmt.Println("Selected Weather:", selectedWeather)

  for _, orbit := range firstCaseOrbits {
    fmt.Println("Enter current max speed for", orbit.Name)
    fmt.Scan(&orbit.CurrentSpeed)
  }

  for _, vehicle := range selectedWeather.AllowedVehicles {
    for _, orb := range firstCaseOrbits {
      estimatedTime := calculateEstimatedTime(sunny, orb, vehicle)
      ov := &OrbitVehicle{Vehicle: vehicle, Orbit: orb, EstimatedTime: estimatedTime}
      fmt.Println("OrbitVehicle: Name:", orb.Name, ", Vehicle:", vehicle.Name, ", Estimated Time:", estimatedTime, " minutes")
      orbitsSpeeds = append(orbitsSpeeds, ov)
    }
  }

  sort.Slice(orbitsSpeeds[:], func(i, j int) bool {
    return orbitsSpeeds[i].EstimatedTime < orbitsSpeeds[j].EstimatedTime
  })
  fmt.Println("Best orbit details:")
  orbSpeed := orbitsSpeeds[0]
  fmt.Println("OrbitVehicle: Name:", orbSpeed.Orbit.Name, ", Vehicle:", orbSpeed.Vehicle.Name, ", Estimated Time:", orbSpeed.EstimatedTime, " minutes")
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

// a func where the source and destination have no intermediary places.
func handleFirstCase(orbits []*Orbit, source *Place, destination *Place) []*Orbit {
  filteredOrbits := []*Orbit{}
  for _, orbit := range orbits {
    if ((orbit.Source == source) && (orbit.Destination == destination)) || ((orbit.Source == destination) && (orbit.Destination == source)) {
      filteredOrbits = append(filteredOrbits, orbit)
    }
  }
  return filteredOrbits
}

// a func which uses dijkstra's algo to find orbits from source to destination.
func handleSecondCase(orbits []*Orbit, source *Place, destination *Place) []*Orbit {
  return []*Orbit{}
}

func getPlace(name string, places []*Place) *Place {
  var place *Place
  for _, plc := range places {
    if plc.Name == name {
      place = plc
      return place
    }
  }
  return place
}
