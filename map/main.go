package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[int]string)

	// colors[10] = "#FFFFFF"

	// delete(colors, 10)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
