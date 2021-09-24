package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMsg := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
			//t.Errorf("got: ", got, "    want: ", want)
		}
	}
	t.Run("saying Hello to people", func(t *testing.T) {
		name := "Chris"
		got := Hello(name)
		want := "Hello, " + name
		assertCorrectMsg(t, got, want)
	})
	t.Run("saying 'Hello, World'", func(t *testing.T) {
		name := "World"
		got := Hello(name)
		want := "Hello, " + name

		assertCorrectMsg(t, got, want)
	})
	t.Run("saying 'Hello, World' when an empty string is given", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMsg(t, got, want)
	})
}
