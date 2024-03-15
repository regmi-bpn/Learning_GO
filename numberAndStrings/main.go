package main

import "fmt"

func main() {

	//strings
	var name string = "bipin"
	var nameOne = "Buddha"
	var nameTwo string

	fmt.Println(name, nameOne, nameTwo)

	name = "treeleaf"
	nameTwo = "newariBhutu"
	fmt.Println(name, nameOne, nameTwo)

	okay := "okay"

	fmt.Println(okay)

	//integers

	var ageOne int = 12
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint = 25

	fmt.Println(numOne, numTwo, numThree)

	//float
	var scoreOne float32 = 23.23
	var scoreTwo float64 = 747384783.34843984
	scoreThree := 12.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
