/*
	Package terrain provides a terrain generator.

	Simplex noise is used to generate the terrain, with multiple layers
	of noise added to create a more interesting result. Tree, Rocks, etc
	are placed on the terrain at deterministic locations.

	The terrain is partioned into Tiles, each tile has a 16x16 grid with
	up to 256 Nodes. An individual Node occupies a 8x8 (world units) plot.
	This makes the Tile 128x128 (world units) with 64*64 Plots.
*/
package terrain

import (
	"realmoforder.com/model/foliage"
	"realmoforder.com/model/mineral"
	"realmoforder.com/model/spatial"
)

// Type of terrain at a particular point.
type Type struct {
	Heat uint8  // Temperature in Celsius.
	Rain uint8  // Precipitation.
	Feet uint16 // Height in feet relative to (waterlevel - 9).
}

// Request terrain modification.
type Request struct {
	Type Type
	Area spatial.Model
}

// Model is responsible for generating the terrain for a world.
type Model struct {
	Seed int64 // deterministic seed for the random number generator.

	Mods []Request // changes made to the terrain.
}

// Slice of terrain and organic cover,
// primarily useful as a visual hint.
type Response struct {
	Terrain []Type          // biome and heightmap.
	Foliage []foliage.Model // organic cover.
	Pebbles []mineral.Model // rocks/stones
}
