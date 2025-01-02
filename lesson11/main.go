package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")
	/* giá trị mybill được sao chép (copy) và truyền vào hàm. */
	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 8.95)
	mybill.addItem("toffee pudding", 4.95)
	mybill.addItem("coffee", 3.25)

	mybill.updateTip(10)
	fmt.Println(mybill.format())
	// go run main.go bill.go
}