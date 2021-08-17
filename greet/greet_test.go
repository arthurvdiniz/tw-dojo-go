package greet

import (
  "testing"
)

func TestGreeting(t *testing.T) {
  t.Run("testing if returns greeting", func(t *testing.T) {
    want := "Greeting from Golang"
    got := Greeting()

    if want != got {
      t.Errorf("\nError:\n want\t-> %s\n got\t-> %s", want, got)
    }
  })
}