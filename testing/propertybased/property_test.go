package propertybased

import (
	"testing"
	"testing/quick"
)

func Mul(a1, a2 int) int {
	return a1 * a2
}

func Test(t *testing.T) {
	f := func(a1, a2 int) bool {
		m := Mul(a1, a2)

		return m == a1*a2
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
