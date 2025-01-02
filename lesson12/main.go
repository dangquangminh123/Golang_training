package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// Cách 1
	// fmt.Print("Create a new bill name:")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	// Cách 2
	name, _ := getInput("Create a new bill name:", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t  -  add tip): ", reader)
	fmt.Println(opt)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}

// go run main.go order.go