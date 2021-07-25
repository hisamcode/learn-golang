package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	colors := make(map[string]string)
	colors["white"] = "ffffff"
	colors["gaga"] = "gaga"

	delete(colors, "gaga")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex color for", color, "is", hex)
	}
}