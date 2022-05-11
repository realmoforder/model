// Package reality provides a data model for a reality.
package reality

import (
	"realmoforder.com/model/citizen"
	"realmoforder.com/model/creator"
	"realmoforder.com/model/critter"
	"realmoforder.com/model/culture"
	"realmoforder.com/model/entropy"
	"realmoforder.com/model/mindset"
	"realmoforder.com/model/network"
	"realmoforder.com/model/shelter"
	"realmoforder.com/model/sundial"
	"realmoforder.com/model/terrain"
	"realmoforder.com/model/trinket"
	"realmoforder.com/model/vehicle"
)

// Model of reality.
type Model struct {
	Name network.Name
	Size uint32

	Time, Tick sundial.Time

	Terrain terrain.Model
	Entropy entropy.Model
	Network network.Model

	Mindsets []mindset.Model
	Trinkets []trinket.Model
	Cultures []culture.Model
	Citizens []citizen.Model
	Critters []critter.Model
	Shelters []shelter.Model
	Vehicles []vehicle.Model
	Creators []creator.Model
}

// Delta describes a change to the Model.
type Delta struct {
	Model uint8
	Index uint16
	Field uint8
	Pause bool // the field is a pathway.Model that will be stopped at 'Value'
	Value interface{}
}

// Request access to the 'Name' reality with the given size,
// amount of entropy and pool of cultures.
type Request struct {
	Name    network.Name
	Size    uint32
	Entropy entropy.Model
	Culture []culture.Model
}
