package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectGreeting := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Given a name, should greet personally", func(t *testing.T) {
		got := Hello("Eugene", "")
		want := "Hello, Eugene!"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("Given a blank name, should show default greeting", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("Should greet a person in Russian", func(t *testing.T) {
		got := Hello("Жека", "Russian")
		want := "Привет, Жека!"
		assertCorrectGreeting(t, got, want)
	})
}
