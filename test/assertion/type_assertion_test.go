package assertion

import (
	"reflect"
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	agtest "github.com/sku0x20/assertgo/test"
)

func typeOf[T any]() reflect.Type {
	return reflect.TypeOf(*new(T))
}

func Test_TypeAssertion_IsType(t *testing.T) {
	mock, s := agtest.NewSink()
	a := assertion.NewType(matcherasserter.New(s, any(42)))
	a.IsType(typeOf[int]())
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}
