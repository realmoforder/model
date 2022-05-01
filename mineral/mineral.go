// Package provides a data model for minerals.
package mineral

import "realmoforder.com/model/spatial"

// Type identifies a mineral type.
type Type struct {
	Rock uint8 // stone
	Iron uint8 // metal
	Gold uint8 // value
	Coal uint8 // power
}

// Model of a mineral deposit.
type Model struct {
	Type Type
	Area spatial.Model
}
