package assertgo

import "testing"

func T(t testing.TB) {
	t.FailNow()
}

func Ts(t testing.TB) {
	t.Fail()
}
