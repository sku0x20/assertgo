package assertion

import (
	"reflect"
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_TypeAssertion_IsType(t *testing.T) {
	mock, s := agtest.NewSink()
	a := assertion.NewType(matcherasserter.New(s, 42))
	a.IsType(reflect.TypeOf(0))
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}
