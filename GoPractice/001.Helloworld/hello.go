package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please enter your name: ")
	var name string
	fmt.Scanln(&name)

	name = strings.TrimSpace(name)
	fmt.Printf("Hi, %s! I'm go!", name)
}
