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
}
