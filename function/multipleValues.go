package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn, sn := getInitials("one piece")
	fmt.Println(fn, sn)

	fn1, sn1 := getInitials("solo leveling")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("naruto")
	fmt.Println(fn2, sn2)

}
