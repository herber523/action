package echo

import "testing"

func TestEcho(t *testing.T) {
	want := "Hello, World!"
	if got := Echo(); got != want {

		t.Errorf("Echo() = %q, want %q", got, want)
	}
}
