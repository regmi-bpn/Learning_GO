package main

import (
	"fmt"
	"sort"
)

func main() {
	//string package example

	//greetings := "hello there friends"

	//fmt.Println(strings.Contains(greetings, "hello"))
	//fmt.Println(strings.ReplaceAll(greetings, "hello", "hi"))
	//fmt.Println(strings.ToUpper(greetings))
	//fmt.Println(strings.Index(greetings, "ll"))
	//fmt.Println(strings.Split(greetings, " "))

	ages := []int{85, 10, 97, 36, 65, 42}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	name := []string{"rohit", "sunil", "rustam"}
	sort.Strings(name)
	fmt.Println(name)

	fmt.Println(sort.SearchStrings(name, "sunil"))
}
