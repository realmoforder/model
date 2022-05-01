// Package shelter provides a data model for shelters.
package shelter

import (
	"realmoforder.com/model/display"
	"realmoforder.com/model/spatial"
)

// Type identifies a shelter type.
type Type int8

const (
	// CitizenCapital, a city-center of significance.
	// Each named city is defined by its CitizenCapital.
	CitizenCapital Type = iota + 1

	// CitizenTrainer, a physical training facility, for
	// citizens to become warriors.
	CitizenTrainer

	// CitizenLibrary, an education facility, for citizens
	// to become scholars. Data refers to the citizen skill.
	CitizenLibrary

	// CitizenTheatre, an entertainment facility, for citizens
	// to be entertained, and to become entertainers.
	CitizenTheatre

	// CitizenStadium, a sporting arena, for citizens to
	// become athletes, and to be entertained. Data refers to
	// the citizen sport.
	CitizenStadium

	// CitizenCouncil, a government structure, for citizens
	// to become politicians.
	CitizenCouncil

	// CitizenObelisk, a spiritual facility, for citizens
	// to adjust their mindset. Data refers to the mindset.
	CitizenObelisk

	// CitizenDisplay, a display piece, for citizens to
	// admire.
	CitizenDisplay

	// CitizenFirepit, a fire to keep nearby citizens
	// warm.
	CitizenFirepit

	// CitizenLatrine is a toilet/bathroom, used to increase the
	// sanitation of citizens.
	CitizenLatrine

	// CitizenAquifer is a well, fountain or spring, that supplies
	// water for citizens to drink.
	CitizenAquifer

	// CitizenFactory, a household, home or house for
	// citizens to spawn from and live in.
	CitizenFactory

	// CuisineFactory, a farm, or food production facility.
	// Data refers to the food being produced.
	CuisineFactory

	// CritterFactory, a nursery, or pen for animals to
	// spawn from and live in. Data refers to the critter
	// type.
	CritterFactory

	// SailingFactory, a shipyard or port, for ships and boats
	// to spawn from and be housed in. Data refers to the ship
	// type.
	SailingFactory

	// ProductFactory, a factory, blacksmith, or forge for tools,
	// weapons, products and props to be invented and constructed.
	// Data refers to the both the product class and type.
	ProductFactory

	// DefenceFactory, a vehicle construction facility, for
	// creating military vehicles & seige weapons. Data refers
	// to the vehicle type.
	DefenceFactory

	// ProductCapital, a marketplace/store for products, tools,
	// weapons and props to be sold and bought.
	ProductCapital

	// ProductStorage, a warehouse/storage facility where raw
	// materials are stored.
	ProductStorage

	// ProductConduit, a trading post and/or waypoint, to be used
	// to mark trading routes.
	ProductConduit

	// DefenceConduit, a defence tower, used to spot foreign
	// soldiers and provide cover and a means for engagement.
	DefenceConduit

	// DefenceBarrier, a wall, used to prevent throughfare of
	// foreign soldiers.
	DefenceBarrier

	// DefenceDoorway, a gate that permits entry of citizens
	// and friendly soliders and prevents entry of foreign
	// soldiers.
	DefenceDoorway

	// DefenceTrainer, a training facility, that trains citizens
	// to become archers.
	DefenceTrainer

	// DefencePillary, a jail/prison, used to hold foreign
	// soldiers and/or citizens who break the law.
	DefencePillary

	// DefencePyramid, a cemetery/tomb, where the deceased
	// are put to rest and remembered.
	DefencePyramid
)

// City identifies a city.
type City uint16

// Name identifies a shelter.
type Name uint16

// Model for a shelter.	A shelter is a building or construct by
// citizens in order to serve a particular utility or purpose.
type Model struct {
	Name Name

	// if true, the shelter should not be visible,
	// or be displayed as rubble.
	Dead bool

	// Type of the shelter.
	Type Type

	Skin display.Name // cultural appearance.

	// City identifies the city this shelter resides
	// within.
	City City

	// Area occupied by the shelter.
	Area spatial.Model

	// Data packs data associated to the Type, for
	// example, if this was a CuisineFactory, the
	// data would identify the food that the shelter
	// produces. See the Type for an indication of
	// what Data refers to.
	Data uint64
}

// Request a shelter to be constructed.
type Request struct {
	Type Type
	Skin display.Name
	Area spatial.Model
}
