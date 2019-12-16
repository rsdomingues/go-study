package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"black": "#FFFFFF",
		"white": "#000000",
	}

	colors["purple"] = "#FF00FF"
	delete(colors, "black")

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, hex)
	}
}
