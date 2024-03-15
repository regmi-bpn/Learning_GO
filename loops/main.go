package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("The value of x is ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("The value of x is ", i)
	}

	names := []string{"tree", "leaf", "tech", "nolo", "gies"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[0])
	}

	for index, value := range names {
		fmt.Printf("the value of index at %v is %v \n", index, value)

	}

}
