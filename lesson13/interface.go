package main

import (
	"fmt"
	"math"
)

// shape interface
// interface trong Go để xử lý nhiều loại hình học khác nhau (hình vuông, hình tròn).
/* 
	+Tính trừu tượng: Interface shape giúp gom các phương thức chung.
	+Dễ mở rộng: Nếu thêm loại hình mới, chỉ cần implement các phương thức area và circumf.
*/
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}