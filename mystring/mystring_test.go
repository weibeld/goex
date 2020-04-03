package mystring

import "testing"

func TestString(t *testing.T) {
	want := "MyString"
	if got := MyString(); got != "MyString" {
		t.Errorf("MyString(): got %s, want %s", got, want)
	}
}
