// Package citizen provides a data model for citizens.
package citizen

import (
	"realmoforder.com/model/creator"
	"realmoforder.com/model/display"
	"realmoforder.com/model/entropy"
	"realmoforder.com/model/foliage"
	"realmoforder.com/model/mindset"
	"realmoforder.com/model/mineral"
	"realmoforder.com/model/pathing"
	"realmoforder.com/model/posture"
	"realmoforder.com/model/product"
	"realmoforder.com/model/shelter"
	"realmoforder.com/model/sundial"
	"realmoforder.com/model/utensil"
	"realmoforder.com/model/vehicle"
)

// Name identifies a citizen.
type Name uint16

// Emotion hints at the citizen's
// well-being.
type Emotion uint8

const (
	Happy Emotion = iota
	Sad
	Angry
	Blessed
	Grief
	Scared
	Bored
	Tired
	Hungry
	Thirsty
	Sleepy
	Worried
	Confused
	Unwell
	Injured
)

// Gear bitfield.
type Gear uint64

// Type of citizen.
type Type int8

// Model of a citizen.
type Model struct {
	Dead bool

	Name Name          // identifies the citizen.
	Type Type          // male or female?
	Star bool          // generated a star?
	Path pathing.Model // animation path.
	City shelter.City  // city citizen is from.
	Home shelter.Name  // home of the citizen.
	Mind mindset.Name  // idealogy/mindset.
	Ship vehicle.Name  // ship the citizen is on.
	Deck shelter.Name  // deck the citizen is on.
	From creator.Name  // creator of this citizen.
	Bias int8

	Cash uint8 // level of wealth.
	Ammo uint8 // level of ammo.

	Good, Evil entropy.Model // personality traits.

	Skin display.Name

	Gear utensil.Type // inventory.

	Task map[posture.Type]sundial.Span // experience amount.
}

// Response is the query of the state of a citizen.
type Response struct {
	Face Emotion // emotional state.

	// resources available to this unit.
	Foliages []foliage.Type
	Minerals []mineral.Type
	Products []product.Type
	Utensils []utensil.Type
}

// Request a citizen to take a certain path.
type Request struct {
	Name Name
	Path pathing.Model
}
