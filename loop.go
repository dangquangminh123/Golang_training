package main
import (
	"fmt"
)

func main() {
	// c1 for
	// x := 0;
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// c2 for
	for i := 0; i <5; i++ {
		fmt.Println("value of x is:", i)
	}

	// exercises loop
	names := []string{"PHP", "LARAVEL", "Nodejs", "Golang"}

	for i :=0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println("The value at index %v is %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}

	fmt.Println(names)

	// Booleans & conditionals
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)
	
	if age < 30 {
		fmt.Println("age is less than 30")
	}else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at position", index)
			// Khi gặp continue, vòng lặp bỏ qua các câu lệnh còn lại và chuyển sang lần lặp kế tiếp.
			continue
		}
		fmt.Printf("the value at position %v is %v \n", index, value)
	}

}