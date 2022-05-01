// Package vehicle provides a data model for vehicles.
package vehicle

import (
	"realmoforder.com/model/display"
	"realmoforder.com/model/pathing"
	"realmoforder.com/model/spatial"
)

// Name identifies a vehicle.
type Name uint16

// Type of vehicle, positive is a ship and magnitude indicates the capacity.
// Negative is a siege vehicle and indicates the range.
type Type int8

// Mode that the vehicle is in.
type Mode int8

const (
	ModeTravel Mode = iota
	ModeAttack
)

// Model of a vehicle.
type Model struct {
	Dead bool

	Name Name
	Type Type

	Mode Mode

	Skin display.Name

	Area spatial.Model
	Path pathing.Model
}

// Request a vehicle to move along a path.
type Request struct {
	Name Name
	Mode Mode
	Path pathing.Model
}
