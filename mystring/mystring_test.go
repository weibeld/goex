package mystring

import "testing"

func TestString(t *testing.T) {
	want := "MyString"
	if got := Get(); got != "MyString" {
		t.Errorf("Get(): got %s, want %s", got, want)
	}
}
