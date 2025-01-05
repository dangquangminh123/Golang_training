package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	/* 
		ReadString là phương thức của bufio.Reader: Đọc dữ liệu đầu vào cho đến khi gặp ký tự kết thúc cụ thể
	*/
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	/* 
		bufio là gói (package): Được sử dụng để đọc và ghi dữ liệu có đệm (buffered I/O).
		bufio.NewReader: Tạo một đối tượng Reader, đọc dữ liệu từ một nguồn cụ thể (ở đây là os.Stdin, tức đầu vào từ bàn phím
		os là gói (package): Cung cấp giao diện để làm việc với hệ điều hành: 
			+Làm việc với tệp tin, thư mục.
			+Xử lý đầu vào/đầu ra tiêu chuẩn (stdin, stdout, stderr).
	*/ 
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
	// fmt.Println(opt)

	switch opt {
	case "a" :
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println(name, price)
		promptOptions(b)
	case "t" :
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added - ", tip)
		promptOptions(b)
	case "s" :
		b.save()
		fmt.Println("you choose to save bill - ", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	/* Chương trình bắt đầu là đầu tiên nhập tên hoá đơn trước,
		Tiếp tục tiếp phần prompOptions để tiếp tục các tuỳ chọn add item, save bill, add tip từ người dùng

	*/
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}

// go run main.go order.go