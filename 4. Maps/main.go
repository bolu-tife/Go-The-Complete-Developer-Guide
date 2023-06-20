package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red": "red",
	// 	"green": "green",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["red"] = "RED"

	// delete(colors, "red")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, val := range c {
		fmt.Println("color", color, "=", val)
	}
}
