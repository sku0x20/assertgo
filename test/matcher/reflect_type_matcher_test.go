package matcher

import (
	"reflect"
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_ReflectTypeMatcher_Match_Pass(t *testing.T) {
	m := matcher.NewReflectTypeMatcher(reflect.TypeOf(0))
	if !m.Match(42) {
		t.Fatal("expected match to be true")
	}
}

func Test_ReflectTypeMatcher_Match_Fail(t *testing.T) {
	m := matcher.NewReflectTypeMatcher(reflect.TypeOf(0))
	if m.Match("hello") {
		t.Fatal("expected match to be false")
	}
}

func Test_ReflectTypeMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewReflectTypeMatcher(reflect.TypeOf(0))
	if m.FailureMsg("hello") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
