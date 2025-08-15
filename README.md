# Golang Benchmarking for Datadog Tag Operations

This suite aims to compare different methodologies for how to efficiently manage/store/udpate tag arrays for
datadog metrics.

### `map[string]string` Reuse Test

We have a function that creates a "base" slice every call then overwrites data to some fields. This tries to
evaluate what is the fastest way to do this.

go test -bench=. -benchmem
goos: linux
goarch: arm64
pkg: github.com/tebriel/slicebench
BenchmarkMapRecreateBase-5       7462564               156.0 ns/op           336 B/op          2 allocs/op
BenchmarkMapClone-5              5761983               206.4 ns/op           336 B/op          2 allocs/op
PASS
ok      github.com/tebriel/slicebench   2.356s
