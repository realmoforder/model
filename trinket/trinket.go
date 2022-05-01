// Package trinket provides a data model for trinkets.
package trinket

import (
	"realmoforder.com/model/spatial"
)

// Type of trinket.
type Type uint8

// Model of a trinket.
type Model struct {
	Type Type
	Area spatial.Model
}
