// Package culture provides a culture model.
package culture

import (
	"realmoforder.com/model/creator"
	"realmoforder.com/model/display"
	"realmoforder.com/model/shelter"
	"realmoforder.com/model/vehicle"
)

// Name identifies a particular culture.
type Name uint8

// Model of a culture.
type Model struct {
	Creators []creator.Name // creators of the culture.
	Trinkets []display.Model
	Shelters map[shelter.Type][]display.Model
	Vehicles map[vehicle.Type][]display.Model
}

// Request a new culture.
type Request struct{}
