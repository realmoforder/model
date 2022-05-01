// Package critter provides a data model for critters/animals.
package critter

import (
	"realmoforder.com/model/pathing"
	"realmoforder.com/model/shelter"
)

// Name idenfifies a critter.
type Name uint16

// Type of critter.
type Type uint8

// Model of a critter.
type Model struct {
	Dead bool

	Name Name  // identity
	Land bool  // or sea
	Male bool  // or female
	Type Type  // class
	Risk uint8 // danger level
	City shelter.City
	Home shelter.Name  // city it was from
	Path pathing.Model // animation
}

// Request a critter to take a certain path.
type Request struct {
	Name Name
	Path pathing.Model
}
