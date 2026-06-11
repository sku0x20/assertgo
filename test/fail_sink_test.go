package test

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
)

func Test_TSink_Fail(t *testing.T) {
	mock := &MockT{}
	sink := pkg.NewTSink(mock)
	sink.Fail("something went wrong")
	if !mock.FatalCalled {
		t.Fatal("expected Fatal to be called")
	}
	if mock.FatalMsg != "something went wrong" {
		t.Fatalf("expected message 'something went wrong', got '%s'", mock.FatalMsg)
	}
}
