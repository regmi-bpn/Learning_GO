package main

import (
	"fmt"
	"math"
)

func sayGreetings(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreetings("prabin")
	sayBye("prabin")
	cycleNames([]string{"okay", "lets", "do", "this"}, sayGreetings)
	cycleNames([]string{"okay", "lets", "do", "this"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("The circle 1 area is %0.1f and circle 2 area is %0.1f ", a1, a2)
}
