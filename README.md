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
BenchmarkMapRecreateBase-5               8694950               135.4 ns/op           336 B/op          2 allocs/op
BenchmarkMapClone-5                      6512355               180.0 ns/op           336 B/op          2 allocs/op
BenchmarkSliceUnroll-5                  22983133                54.44 ns/op          256 B/op          3 allocs/op
BenchmarkSliceLiteralItems-5            24743875                40.69 ns/op          192 B/op          2 allocs/op
BenchmarkRuntimeSliceExpansion-5        11247483               109.6 ns/op           560 B/op          3 allocs/op
BenchmarkSlicePreallocation-5           17100126                76.78 ns/op          480 B/op          1 allocs/op
PASS
ok      github.com/tebriel/slicebench   7.158s
```
