package main

import "fmt"

func main() {

	//var ages [3]int = [3]int{20, 10, 11}
	var ages = [3]int{20, 10, 11}
	name := [3]string{"rohit", "sunil", "rustam"}
	name[0] = "prabin"
	fmt.Println(ages, len(ages), name, len(name))

	//slices
	var scores = []int{10, 20, 30}
	scores[2] = 25
	scores = append(scores, 45)
	fmt.Println(scores)

	//slice ranges
	rangeOne := name[1:3]
	rangeTwo := name[2:]
	rangeThree := name[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
