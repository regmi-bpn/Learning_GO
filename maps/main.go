package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":    40.56,
		"pie":     50.66,
		"pudding": 30.65,
		"souffle": 90.32,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// integer as key types
	phoneBook := map[int]string{
		12323223: "marie",
		34343433: "louise",
		94343433: "henry",
		74343433: "saud",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[74343433])

	phoneBook[74343433] = "cavill"
	fmt.Println("The changed name is:", phoneBook[74343433])
}
