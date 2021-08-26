package greet

import "fmt"

func Greeting(name string) (output string) {
  if name == "" {
    name = "my friend"
  }

  output = fmt.Sprintf("Hello, %s.", name)
  return
}