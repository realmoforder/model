// Package entropy provides a data model for entropy.
package entropy

import (
	"realmoforder.com/model/aqueous"
	"realmoforder.com/model/critter"
	"realmoforder.com/model/foliage"
	"realmoforder.com/model/mineral"
	"realmoforder.com/model/product"
	"realmoforder.com/model/shelter"
	"realmoforder.com/model/trinket"
	"realmoforder.com/model/utensil"
	"realmoforder.com/model/vehicle"
)

// Model of complexity.
type Model struct {
	Entropy uint8 // which field is important?

	Mineral mineral.Type
	Aqueous aqueous.Type
	Foliage foliage.Type
	Shelter shelter.Type
	Vehicle vehicle.Name
	Critter critter.Type
	Product product.Name
	Utensil utensil.Type
	Display trinket.Type
}
