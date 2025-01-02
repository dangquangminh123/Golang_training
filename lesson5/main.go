package main

import (
	"fmt"
	// "strings"
)

// Var thì gắn giá trị cho biến và kiểu của biến dùng trong và ngoài hàm đều được!

// Hàm main là hàm để thực thi đầu tiên!
func main() {
	var nameOne string = "PHP"
	var nameTwo = "LARAVEL"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Golang"
	nameThree = "Backend"
	fmt.Println(nameOne, nameTwo, nameThree)

	// Dấu : này tự gắn giá trị và kiểu của biến nhưng chỉ dùng trong hàm! 
	nameFour := "Nodejs" 
	fmt.Println(nameFour)

	// int
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	ageFour := "40"

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bit & memory
	var numOne int = 20
	var numTwo int8 = -128
	var numThree uint8 = 128
	var numFour uint16 = 256

	fmt.Println(numOne, numTwo, numThree, numFour)
	// Float số thực 
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 489290392003990.6
	scoreThree := 21.54

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	//Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello developer!")
	fmt.Println("goodbye developer!")

	fmt.Println("My age is", ageOne, "And my name is", nameOne);

	// Printf (formatted strings)
	/* %v Hiển thị giá trị của biến ở dạng mặc định, tùy thuộc vào kiểu dữ liệu.*/ 	
	fmt.Printf("my age is %v and my name %v \n", ageTwo, nameTwo);
	/* %q Hiển thị giá trị của chuỗi (string) kèm theo dấu nháy kép.*/ 
	fmt.Printf("my age is %q and my name %q \n", ageFour, nameThree);

	/* %T Hiển thị kiểu dữ liệu của giá trị.*/ 
	fmt.Printf("age is of type %T \n", ageThree);

	/* %f Hiển thị số thực với mặc định 6 chữ số thập phân. */ 
	fmt.Printf("you scored %f points! \n", 255.55);
	fmt.Printf("you scored %0.1f points \n", 255.55);

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", ageFour, nameFour)
	fmt.Println("the saved string is", str)
}

