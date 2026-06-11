package assertgo

import "testing"

func T(t *testing.T) {
	t.Fail()
}

func Ts(t *testing.T) {
	t.Fail()
}
