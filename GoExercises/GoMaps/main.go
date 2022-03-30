package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#f0000",
		"blue":  "#f0010",
		"white": "#fffff",
	}

	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["white"] = "#fffff"
	// delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for color", color, "is", hex)
	}
}
