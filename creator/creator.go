// Package creator provides a data model for creators.
package creator

import (
	"realmoforder.com/model/shelter"
)

// Name identifies a creator.
type Name uint8

// User name.
type User string

// Model of a creator/player/user.
type Model struct {
	Name Name

	User User   // name of the user
	Dust uint64 // score / points

	Allies []Name
	Cities []shelter.City
}

// Request an alliance with another
// creator.
type Request struct {
	Name Name
	Ally bool
}
