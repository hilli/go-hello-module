package hello

import (
	"fmt"
	"strings"
)

// SayHello just says Hello
func SayHello() {
	fmt.Println("Hello")
}

// SayHelloTo says hello to all present
func SayHelloTo(names ...string) {
	fmt.Println("Hello", strings.Join(names, ", "))
}
