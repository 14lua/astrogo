package orbits

import (
	"math"
	"github.com/14lua/astronomy-lib/globals"
)

func CircularVelocity(massKg1, massKg2, radiusFromBarycenterMeters float64) (orbitalVelocity float64) {
	// orbital velocity = v
	M := massKg1 + massKg2 // total mass of system
	orbitalVelocity = math.Sqrt((globals.G * M) / radiusFromBarycenterMeters)
	return
}
