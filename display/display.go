// Package display provides a display model.
package display

import (
	"realmoforder.com/model/spatial"
)

// Name refers to something that can be displayed.
type Name uint16

// Model that can be displayed.
type Model struct {
	Name Name
	Path string // globally unique path to the asset.
	Area spatial.Model
	Deck []spatial.Model // decks available to stand on.
}
