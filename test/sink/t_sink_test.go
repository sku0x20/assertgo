package sink

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/sink"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_TSink_Fail(t *testing.T) {
	mock := &agtest.MockT{}
	s := sink.NewTSink(mock)
	s.Fail("something went wrong")
	if !mock.FatalCalled {
		t.Fatal("expected Fatal to be called")
	}
	if mock.FatalMsg != "something went wrong" {
		t.Fatalf("expected message 'something went wrong', got '%s'", mock.FatalMsg)
	}
}
