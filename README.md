# assertgo

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

`T(t)` returns an `AssertionSuite`. Call `Assert(value)` to get an `AnyAssertion`, then chain assertion methods.

## Assertions

| Method | Description |
|---|---|
| `IsEqualTo(other)` | Fails if value does not equal other |
