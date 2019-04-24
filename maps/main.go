package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Color: %v; Hex Code: %v\n", color, hex)
	}
}
