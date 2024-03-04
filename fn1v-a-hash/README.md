# FNV-1 Hash in Go

This part was inspired by the blog 1 billion records challenge [Go-1BRC](https://benhoyt.com/writings/go-1brc/) here. So there is a last solution part in the optimization phase for rewrite a hash table instead of using `map` in Golang, so it very interesting for me.

This folder is used to explore it, hopefully it make you feel fun.

References:
- [FN1V Wikipedia](https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV-1a_hash)
- [Github Counter](https://github.com/benhoyt/counter/blob/master/counter.go)

Benchmark test:
```bash
goos: darwin
goarch: arm64
pkg: example.com/m/v2
BenchmarkMostlyUniqueCounter-8                     13828             83466 ns/op
BenchmarkNonUniqueCounter-8                        16611             75857 ns/op
BenchmarkMostlyUniqueMapBytes-8                     3159            368908 ns/op
BenchmarkNonUniqueMapBytes-8                        5348            216863 ns/op
BenchmarkMostlyUniqueMapPointerBytes-8              4516            269789 ns/op
BenchmarkNonUniqueMapPointerBytes-8                10000            108574 ns/op
BenchmarkMostlyUniqueMapString-8                    4698            246280 ns/op
BenchmarkNonUniqueMapString-8                      13102             92328 ns/op
PASS
ok      example.com/m/v2        13.727s
```

Please treat this code as research work only, do not using it in production.