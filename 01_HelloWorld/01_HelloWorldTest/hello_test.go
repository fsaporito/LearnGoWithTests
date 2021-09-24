package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, Wrld"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
		//t.Errorf("got: ", got, "    want: ", want)
	}
}
