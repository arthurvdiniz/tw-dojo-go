package greet

import (
	"fmt"
	"strings"
)

func Greeting(names ...string) string {
  greeting := ""

	if len(names) == 1 && names[0] == "" {
		greeting = "my friend"
	}

	for _, v := range names {
		
	}

	if names == strings.ToUpper(names) {
		return fmt.Sprintf("HELLO, %s!", names)
	}

	return fmt.Sprintf("Hello, %s.", names)
}
