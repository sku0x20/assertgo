# AssertGo

A simple assertion library for Go.

## Installation

```sh
go get github.com/sku0x20/assertgo
```

## Usage

```go
import . "github.com/sku0x20/assertgo"

func TestSomething(t *testing.T) {
    T(t).Assert("hello").IsEqualTo("hello")
}
```

