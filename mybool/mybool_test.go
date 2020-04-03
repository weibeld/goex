package mybool

import "testing"

func TestBool(t *testing.T) {
	want := true
	if got := Get(); got != want {
		t.Errorf("Get(): got %t, want %t", got, want)
	}
}
