package main

import "fmt"

func updateName(name string) string {
	name = "wedge"
	return name
}

func updateMenu(menu map[string]float64) {
	menu["coffee"] = 20.99
}

func main() {
	//group A types --> strings, ints, bools, floats, array, struct

	name := "tifa"

	name = updateName(name)
	fmt.Println(name)

	//group B type -->slices, maps, functions

	menu := map[string]float64{
		"soup":    40.56,
		"pie":     50.66,
		"pudding": 30.65,
		"souffle": 90.32,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
