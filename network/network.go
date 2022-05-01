// Package network provides a data model for economic networking.
package network

import (
	"realmoforder.com/model/citizen"
	"realmoforder.com/model/creator"
	"realmoforder.com/model/foliage"
	"realmoforder.com/model/mineral"
	"realmoforder.com/model/product"
	"realmoforder.com/model/trinket"
	"realmoforder.com/model/utensil"
)

// Name identifies a network reachable reality.
type Name string

// Request structure for tradeable types.
type Request[Type any] struct {
	With citizen.Name // citizen who is performing the transfer.
	Mode int8         // negative for export, positive for import.
	Type Type         // resource type
	Name Name         // destination name.
	User creator.User // destination user.
}

type Model struct {
	Trinkets []Request[trinket.Type]
	Utensils []Request[utensil.Type]
	Foliages []Request[foliage.Type]
	Minerals []Request[mineral.Type]
	Products []Request[product.Type]
}
