# AssertGo

A simple assertion library for Go.

## Installation

```sh
go get github.com/sku0x20/assertgo
```

## Usage

```go
import "github.com/sku0x20/assertgo"

func TestSomething(t *testing.T) {
    T(t).Assert("hello").EqualTo("hello")
    T(t).Assert(obj).SameAs(obj)
    T(t).Assert(nil).IsNil()
}
```

### Negation

Chain `Not()` before any assertion to invert it:

```go
T(t).Assert("hello").Not().EqualTo("world")
T(t).Assert(obj).Not().IsNil()
```

### Custom Matchers

For advanced cases, pass a `Matcher[V]` directly:

```go
T(t).Assert("hello").Matches(myMatcher)
```

A matcher implements two methods:

```go
type Matcher[V any] interface {
    Match(value V) bool
    FailureMsg(value V) string
}
```

For example, the built-in `NotMatcher` wraps another matcher and inverts it:

```go
notMatcher := matcher.NewNotMatcher(matcher.NewDeepEqualMatcher("hello"))
T(t).Assert("world").Matches(notMatcher)
```
