// Package pathing provides a data model for pathing.
package pathing

import (
	"realmoforder.com/model/posture"
	"realmoforder.com/model/sundial"
)

// Model of a pathway, a sequence of postures.
type Model struct {
	Loop bool   // does the end of the path connect back to the beginning?
	Pace uint16 // how fast does the path progress.

	Time sundial.Time  // offset from the epoch.
	Stop sundial.Time  // when does the path stop / become disrupted.
	From posture.Model // start of the path.

	Step []posture.Model // steps in the path.
}
