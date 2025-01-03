package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n);
}

func sayBye(n string) {
	fmt.Printf("GoodBye %v \n", n);
}

/* 
	range là một từ khóa trong Go, được sử dụng để lặp qua các phần tử của:
	Mảng (array), Slice, Map, Kênh (channel), Chuỗi (string).
*/ 
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("PHP")
	// sayBye("LARAVEL")
	// sayGreeting("Nodejs")

	cycleNames([]string{"Laravel", "Nodejs", "Golang"}, sayGreeting)
	cycleNames([]string{"PHP", "Nest", "Golang"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println("Area Circle is: \n", a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}