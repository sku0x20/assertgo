package e2e

import "testing"
import . "github.com/sku0x20/assertgo"

func Test_FailHard(t *testing.T) {
	T(t)
}

func Test_FailSoft(t *testing.T) {
	Ts(t)
	if !t.Failed() {
		t.Fatal("should have failed")
	}
}
