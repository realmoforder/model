// Package posture provides data model for citizen postures.
package posture

import (
	"realmoforder.com/model/spatial"
	"realmoforder.com/model/sundial"
)

// Type identifies a posture.
type Type uint8

const (
	// Idle, unit is not doing anything
	// in particular.
	Idle Type = iota

	// Reading, the citizen is reading a
	// book.
	Reading

	// Exercising, the citizen is physically
	// exercising.
	Exercising

	// Acting, the citizen is acting in a
	// theatrical fashion.
	Acting

	// Competing, the citizen is competing in
	// a sport. Data refers to the sport.
	Competing

	// Speaking, the citizen is speaking to a
	// crowd.
	Speaking

	// Praying, the citizen is praying to a
	// deity.
	Praying

	// Building, the citizen is building a
	// structure. Data refers to the structure.
	Building

	// Crafting, the citizen is crafting a
	// product. Data refers to the product.
	Crafting

	// Farming, the citizen is farming food.
	// Data refers to the food.
	Farming

	// Sailing, the citizen is sailing a ship.
	Sailing

	// Preaching, the citizen is preaching to
	// a congregation.
	Preaching

	// Dancing, the citizen is dancing.
	Dancing

	// Eating, the citizen is eating.
	Eating

	// Drinking, the citizen is drinking.
	Drinking

	// Shopping, the citizen is shopping.
	Shopping

	// Fighting, the citizen is fighting.
	Fighting

	// Shooting, the citizen is firing a weapon.
	Shooting

	// Gathering, the citizen is gathering
	// food.
	Gathering

	// Chopping, the citizen is chopping
	// down a tree.
	Chopping

	// Mining, the citizen is mining ore.
	Mining

	// Keeping, the citizen is taking care
	// of an animal. Data refers to the animal.
	Keeping

	// Digging, the citizen is digging a
	// hole.
	Digging

	// Fishing, the citizen is fishing.
	Fishing

	// Trading, the citizen is trading.
	Trading

	// Hunting, the citizen is hunting.
	// Data refers to the animal.
	Hunting

	// Exploring, the citizen is exploring.
	Exploring

	// Performing, the citizen is playing
	// music. Data refers to the instrument.
	Performing

	// Coughing, the citizen is sick/coughing
	// and/or sneezing.
	Coughing

	// Sitting, the citizen is sitting.
	Sitting

	// Sleeping, the citizen is sleeping.
	Sleeping

	// Drowning, the citizen is drowning.
	Drowning

	// Rioting, the citizen is angry, causing
	// trouble.
	Rioting

	// Entertaining, the citizen is juggling and/or
	// cartwheeling. Being entertaining.
	Entertaining

	// Diving, the citizen is diving underwater.
	Diving

	// Cleaning, the citizen is cleaning.
	Cleaning

	// Watering, the citizen is pouring water
	// onto something.
	Watering

	// Planting, the citizen is planting seeds.
	Planting
)

// Model describes a posture, the measurement of an
// activity taking place in a particular area for
// a particular amount of time.
type Model struct {
	Type Type          // type of posturing.
	Data uint8         // data specific to the type.
	Area spatial.Model // area of the posture.
	Span sundial.Span  // duration of the posture.
}
