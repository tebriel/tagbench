# Golang Benchmarking for Datadog Tag Operations

This suite aims to compare different methodologies for how to efficiently manage/store/udpate tag arrays for
datadog metrics.

### `map[string]string` Reuse Test

We have a function that creates a "base" slice every call then overwrites data to some fields. This tries to
evaluate what is the fastest way to do this.

### `[]string` Unroll vs Append Test

We have this pattern of creating a slice and then immediately using it only once as args to `append` with the
`...` operator.

```
go test -bench=. -benchmem
goos: linux
goarch: arm64
pkg: github.com/tebriel/slicebench
BenchmarkMapRecreateBase-5               8726722               131.4 ns/op           336 B/op          2 allocs/op
BenchmarkMapClone-5                      6725186               178.1 ns/op           336 B/op          2 allocs/op
BenchmarkSliceUnroll-5                  23199984                51.79 ns/op          256 B/op          3 allocs/op
BenchmarkSliceLiteralItems-5            32316963                38.04 ns/op          192 B/op          2 allocs/op
BenchmarkSliceRuntimeExpansion-5        11562579               105.4 ns/op           560 B/op          3 allocs/op
BenchmarkSlicePreallocation-5           16364073                71.03 ns/op          480 B/op          1 allocs/op
BenchmarkSliceCreateVariadic-5          42497125                28.42 ns/op          160 B/op          1 allocs/op
BenchmarkSliceCreateWithCopy-5          42560808                28.35 ns/op          160 B/op          1 allocs/op
PASS
ok      github.com/tebriel/slicebench   9.576s
```
