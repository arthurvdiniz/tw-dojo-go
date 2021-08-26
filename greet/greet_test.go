package greet

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	t.Run("test should return a name", func(t *testing.T) {
		want := "Hello, Bob."
		got := Greeting("Bob")

		if want != got {
			t.Errorf("\nError:\n want\t-> %s\n got\t-> %s", want, got)
		}
	})

	t.Run("test should return generic greeting when name is null", func(t *testing.T) {
		want := "Hello, my friend."
		got := Greeting("")

		if want != got {
			t.Errorf("\nError:\n want\t-> %s\n got\t-> %s", want, got)
		}
	})

	t.Run("test should return screaming greeting when name is uppercase", func(t *testing.T) {
		want := "HELLO, JERRY!"
		got := Greeting("JERRY")

		if want != got {
			t.Errorf("\nError:\n want\t-> %s\n got\t-> %s", want, got)
		}
	})

	t.Run("test should return greeting for multiple names", func(t *testing.T) {
		want := "Hello, Jill and Jane."
		got := Greeting("Jill", "Jane")

		if want != got {
			t.Errorf("\nError:\n want\t-> %s\n got\t-> %s", want, got)
		}
	})
}
