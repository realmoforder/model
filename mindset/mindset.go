// Package mindset provides a data model for mindsets.
package mindset

import (
	"realmoforder.com/model/creator"
	"realmoforder.com/model/entropy"
)

// Name identifies a mindset.
type Name uint16

// Model of a mindset.
type Model struct {
	Name Name
	From creator.Name

	Good, Evil entropy.Model
}
