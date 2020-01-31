package main

import "testing"

// multiple test runs further optimized in V3

func TestHello(t *testing.T) {

	t.Run("saying hello to mama", func(t *testing.T) {
		got := Hello("mama")
		want := "Hello, mama"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("blank hello", func(t *testing.T) {
		got := Hello(" ")
		want := "Hello,  "

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
