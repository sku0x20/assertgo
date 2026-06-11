# AssertGo

A simple assertion library for Go.

## Installation

```sh
go get github.com/sku0x20/assertgo
```

## Usage

```go
import (
    . "github.com/sku0x20/assertgo"
    "github.com/sku0x20/assertgo/pkg/matcher"
)

func TestSomething(t *testing.T) {
    T(t).Assert("hello").EqualTo("hello")
    T(t).Assert("hello").Not().EqualTo("world")

    // For advanced cases, pass a Matcher[V] directly
    T(t).Assert("world").Matches(matcher.NewNotMatcher[any](matcher.NewDeepEqualMatcher[any]("hello")))
}
```
