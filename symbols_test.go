package symbols

import "testing"

func TestZero(t *testing.T) {
	z := Symbol{}
	x := Make(z.String())
	if x != z {
		t.Error("zero symbol z != Make(z.String())")
	}
}

func TestEqual(t *testing.T) {
	a := make([]byte, 3)
	b := make([]byte, 3)
	copy(a, "foo")
	copy(b, "foo")
	sa := Make(string(a))
	sb := Make(string(b))
	if sa != sb {
		t.Error("two symbols named foo are not equal")
	}
}

func TestNotEqual(t *testing.T) {
	a := Make("foo")
	b := Make("bar")
	if a == b {
		t.Error(`Make("foo") == Make("bar")`)
	}
}
