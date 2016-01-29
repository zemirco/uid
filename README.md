
# uid

[![Build Status](https://travis-ci.org/zemirco/uid.svg)](https://travis-ci.org/zemirco/uid)
[![GoDoc](https://godoc.org/github.com/zemirco/uid?status.svg)](https://godoc.org/github.com/zemirco/uid)

Generate URL safe strings.

## Example

```go
package main

import "github.com/zemirco/uid"

func main() {
  id := uid.New(10)
  fmt.Println(id)
  // 9BZ1sApAX4
}
```

## Test

`go test`

## Benchmark

```bash
BenchmarkNew1   20000000  104 ns/op
BenchmarkNew5   20000000  110 ns/op
BenchmarkNew10  10000000  120 ns/op
BenchmarkNew20  10000000  155 ns/op
BenchmarkNew50   5000000  252 ns/op
```

## License

MIT
