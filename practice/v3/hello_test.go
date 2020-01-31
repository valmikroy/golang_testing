package main

import "testing"

// multiple test with shared code
// -t.Helper() inside function is doing some trick

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to mama", func(t *testing.T) {
		got := Hello("mama")
		want := "Hello, mama"

		assertCorrectMessage(t, got, want)

	})
	t.Run("blank hello", func(t *testing.T) {
		got := Hello(" ")
		want := "Hello,  "

		assertCorrectMessage(t, got, want)
	})
}
