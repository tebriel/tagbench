package bench

import (
	"fmt"
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

// BenchmarkSliceRuntimeExpansion creates the slice and appends items to it, causing runtime expansions
func BenchmarkSliceRuntimeExpansion(b *testing.B) {
	var r, data []string
	for i := range [12]struct{}{} {
		data = append(data, fmt.Sprintf("item-%d", i))
	}

	for i := 0; b.Loop(); i++ {
		r = []string{
			"one",
			"two",
			"three",
			"four",
			"five",
		}
		r = append(r, "five.five")
		// Add twelve more items to the slice
		r = append(r, data...)
		r = append(r, "six")
		r = append(r, "seven")
	}
	sliceResult = r
}

// BenchmarkSliceRuntimeExpansion sets capacity at creation of the slice to avoid unnecessary pauses for
// expansion
func BenchmarkSlicePreallocation(b *testing.B) {
	var r, data []string
	for i := range [12]struct{}{} {
		data = append(data, fmt.Sprintf("item-%d", i))
	}

	for i := 0; b.Loop(); i++ {
		// Pick a larger than expected capacity but not significantly so
		r = make([]string, 0, 30)
		r = append(r,
			"one",
			"two",
			"three",
			"four",
			"five",
		)
		r = append(r, "five.five")
		// Add twelve more items to the slice
		r = append(r, data...)
		r = append(r, "six")
		r = append(r, "seven")
	}
	sliceResult = r
}
