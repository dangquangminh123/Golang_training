package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("Developer Blockchain")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("Developer PHP LARAVEL")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("Developer PHP LARAVEL")
	fmt.Println(fn3, sn3)
}