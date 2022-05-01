// Package foliage provides a data model for organic resources.
package foliage

import "realmoforder.com/model/spatial"

// Type of organic resource.
type Type struct {
	Food uint8
	Tree uint8
	Bush uint8
	Farm uint8
}

// Model of an organic resource.
type Model struct {
	Type Type
	Area spatial.Model
}
