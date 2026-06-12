[![CI](https://github.com/sku0x20/assertgo/actions/workflows/ci.yaml/badge.svg)](https://github.com/sku0x20/assertgo/actions/workflows/ci.yaml)

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
    T(t).Assert("hello").EqualTo("hello")
    T(t).Assert("hello").Not().EqualTo("world")

    T(t).AssertBool(true).IsTrue()
    T(t).AssertBool(false).Not().IsTrue()

    T(t).AssertInt(10).GreaterThan(5)
    T(t).AssertInt(10).InRange(1, 20)
    T(t).AssertFloat(3.14).LessThan(4.0)

    T(t).AssertString("hello").Contains("ell")
    T(t).AssertString("hello").HasPrefix("hel")
    T(t).AssertString("hello").HasSuffix("llo")
    T(t).AssertString("hello123").MatchesRegex(`^\w+\d+$`)

    T(t).AssertSlice([]any{1, 2, 3}).HasLength(3)
    T(t).AssertSlice([]any{1, 2, 3}).ContainsAll([]any{1, 3})
    T(t).AssertSlice([]any{1, 2, 3}).Not().ContainsAll([]any{4, 5})
    T(t).AssertSlice([]any{1, 2, 3}).IsUnique()

    T(t).AssertMap(map[any]any{"a": 1}).HasLength(1)
    T(t).AssertMap(map[any]any{"a": 1, "b": 2}).ContainsAllEntries(map[any]any{"a": 1})

    T(t).AssertType(42).IsType(reflect.TypeOf(0))
}
```

## Custom Comparator

```go
T(t).Assert(myStruct).WithComparator(myComparator).EqualTo(expected)
```

## Custom Matcher

```go
import "github.com/sku0x20/assertgo/pkg/matcher"

T(t).Assert("hello").Matches(matcher.NewNotMatcher[any](matcher.NewNilMatcher[any]()))
```
