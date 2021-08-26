package greet

import (
	"fmt"
	"strings"
)

func Greeting(name ...string) string {
	if name == "" {
		name = "my friend"
	}

	if name == strings.ToUpper(name) {
		return fmt.Sprintf("HELLO, %s!", name)
	}

	return fmt.Sprintf("Hello, %s.", name)
}
