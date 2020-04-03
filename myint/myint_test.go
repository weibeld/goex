package myint

import "testing"

func TestInt(t *testing.T) {
	want := 42
	if got := MyInt(); got != want {
		t.Errorf("MyInt(): got %d, want %d", got, want)
	}
}
