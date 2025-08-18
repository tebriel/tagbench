package bench

import (
	"testing"
)

var sliceResult []string

// BenchmarkSliceRecreateBase creates the "data" and then uses ... to unroll it
func BenchmarkSliceUnroll(b *testing.B) {
	var r, data []string
	for i := 0; b.Loop(); i++ {
		r = []string{
			"one",
			"two",
			"three",
			"four",
		}
		// make a slice then immediately unroll it
		data = []string{
			"twotwo",
			"fourfour",
			"sixsix",
			"eighteight",
		}
		r = append(r, data...)
	}
	sliceResult = r
}

// BenchmarkSliceReuseBase creates the "data" and then uses append to add it
func BenchmarkSliceLiteralItems(b *testing.B) {
	var r []string
	for i := 0; b.Loop(); i++ {
		r = []string{
			"one",
			"two",
			"three",
			"four",
		}
		r = append(r, "twotwo",
			"fourfour",
			"sixsix",
			"eighteight",
		)
	}
	sliceResult = r
}
