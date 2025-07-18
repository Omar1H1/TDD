package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Chris")
		want := "Hello, Chris"

	  assertCorrectMessage(t, got, want)	

	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {

		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)

	})
}

// Helpers allows us to reduce code duplication 
// t.Helper() so the compailer show the line in the original function and not the helper
func assertCorrectMessage(t testing.TB, got, want string) {

	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
