
# uid

[![Build Status](https://travis-ci.org/zemirco/uid.svg)](https://travis-ci.org/zemirco/uid)
[![GoDoc](https://godoc.org/github.com/zemirco/uid?status.svg)](https://godoc.org/github.com/zemirco/uid)

Generate URL safe strings.

## Example

```go
package main

import "github.com/zemirco/uid"

func main() {
  id, err := uid.Gen(10)
  if err != nil {
    panic(err)
  }
  fmt.Println(id)
  // 9BZ1sApAX4
}
```

## Test

`go test`

## License

MIT
