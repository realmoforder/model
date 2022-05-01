// Package spatial provides a data model for 2D space.
package spatial

// Model describes an area of space, centered at the
// given X, Y coordinates with a size corresponding
// to the given angle, radius and sqaured height.
type Model struct {
	X, Y int32

	Angle float32

	Radius uint16
	Height uint16
}
