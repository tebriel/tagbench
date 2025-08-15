// Package bench for benchmarking different tag concatenation operations
package bench

import (
	"maps"
	"testing"
)

var result map[string]string

// BenchmarkMapRecreateBase creates the base every time and then modifies it
func BenchmarkMapRecreateBase(b *testing.B) {
	var r, data map[string]string

	data = map[string]string{
		"b": "twotwo",
		"d": "fourfour",
	}

	for i := 0; b.Loop(); i++ {
		r = map[string]string{
			"a": "one",
			"b": "two",
			"c": "three",
			"d": "four",
		}
		for key := range data {
			r[key] = data[key]
		}
	}
	result = r
}

// BenchmarkMapClone shallow clones the base map and then modifies it with data
func BenchmarkMapClone(b *testing.B) {
	var base, data, r map[string]string
	base = map[string]string{
		"a": "one",
		"b": "two",
		"c": "three",
		"d": "four",
	}
	data = map[string]string{
		"b": "twotwo",
		"d": "fourfour",
	}
	for i := 0; b.Loop(); i++ {
		r = maps.Clone(base)
		for key := range data {
			r[key] = data[key]
		}
	}
	result = r
}
