package main

import "fmt"

func main() {

	age := 45

	if age < 30 {
		fmt.Println("The age is less than 45")
	} else if age < 40 {
		fmt.Println("The age is less than 40")
	} else {
		fmt.Println("The age is not less than 45")
	}

	names := []string{"tree", "leaf", "techno", "logies"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos ", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at pos ", index)
			break
		}

		fmt.Printf("The value at pos %v is %v \n", index, value)
	}
}
