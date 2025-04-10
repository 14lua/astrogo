package orbits

import (
	"math"
	"testing"
	"github.com/14lua/astrogo/globals"
)

func TestCircularVelocity(t *testing.T) {
	t.Run("within 0.5%, no sattelite mass", func(t *testing.T) {
		altitudeAboveSurfaceMeters := 2000.0 * 1000.0 // 2000km
		got := CircularVelocity(globals.EarthMassKg, 0.0, (globals.EarthRadiusMeters + altitudeAboveSurfaceMeters))
		want := 6900.0 // in m/s
		deviation := (math.Abs(got - want) / want) * 100
		if deviation >= 0.5 {
			t.Errorf("got %f want %f; not within 0.5 percent", got, want)
		}
	})
}
