package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code of ", color, " is ", hex)
	}
}
