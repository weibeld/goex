package mybool

import "testing"

func TestBool(t *testing.T) {
	want := true
	if got := MyBool(); got != want {
		t.Errorf("MyBool(): got %t, want %t", got, want)
	}
}
