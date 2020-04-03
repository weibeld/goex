package myint

import "testing"

func TestInt(t *testing.T) {
	want := 42
	if got := Get(); got != want {
		t.Errorf("Get(): got %d, want %d", got, want)
	}
}
