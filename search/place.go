package search

import (
	"errors"

	"github.com/prasadsurase/gk-traffic-problem-3/structs"
)

// Place : search a place by string among the list of provided places
func Place(name string, places []*structs.Place) (*structs.Place, error) {
	for _, plc := range places {
		if plc.Name == name {
			return plc, nil
		}
	}
	return nil, errors.New("Location not found")
}
