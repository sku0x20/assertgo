package assertgo

import "testing"

func T(t testing.TB) {
	t.Fail()
}

func Ts(t testing.TB) {
	t.Fail()
}
