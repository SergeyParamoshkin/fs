# fs

[origininal code and note](https://www.gopherguides.com/articles/golang-1.16-io-fs-improve-test-performance)

[https://golang.org/pkg/io/fs/](https://golang.org/pkg/io/fs/)

[github issue Russ Cox](https://github.com/golang/go/issues/41190)

[File System Interfaces for Go — Draft Design](https://go.googlesource.com/proposal/+/master/design/draft-iofs.md)
```golang
go test  -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out    
goos: darwin
goarch: amd64
pkg: github.com/SergeyParamoshkin/fs
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkGoFilesJIT-16                      2673            444149 ns/op            6287 B/op         57 allocs/op
BenchmarkGoFilesFS-16                     634747              1856 ns/op             976 B/op         33 allocs/op
BenchmarkGoFilesExistingFiles-16           18918             63722 ns/op            3952 B/op         39 allocs/op
````