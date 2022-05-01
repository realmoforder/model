// Package utensil provides a data model for utensils.
package utensil

type Type uint64

const (
	Axe Type = 0b11 << (iota * 2)
	Pickaxe
	Shovel
	Hoe
	Sickle
	FishingRod
	Bucket
	Knife
	Fork
	Wand
	Clippers
	Basket
	Hammer

	Hat
	Outfit
	Shoes
	Gloves

	Helmet
	Armour
	Sword
	Spear
	Club
	Sling
	Bow
	Shield

	Instrument

	Cape
	Idol
	Potion
	Aura
	Appendage
	Torch
)
