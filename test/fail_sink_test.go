package test

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
)

func Test_FailSink_Fail(t *testing.T) {
	mock := &mockT{}
	sink := pkg.NewFailSink(mock)
	sink.Fail("something went wrong")
	if !mock.fatalCalled {
		t.Fatal("expected Fatal to be called")
	}
	if mock.fatalMsg != "something went wrong" {
		t.Fatalf("expected message 'something went wrong', got '%s'", mock.fatalMsg)
	}
}
