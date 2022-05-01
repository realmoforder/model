// Package product provides a product data model.
package product

import (
	"realmoforder.com/model/aqueous"
	"realmoforder.com/model/critter"
	"realmoforder.com/model/foliage"
	"realmoforder.com/model/mineral"
	"realmoforder.com/model/terrain"
)

// Name identifies a refined product.
type Name uint32

// Type of product, defined by identifying
// the raw material it is derived from.
type Type struct {
	Foliage foliage.Type
	Critter critter.Type
	Mineral mineral.Type
	Aqueous aqueous.Type
	Terrain terrain.Type
}
