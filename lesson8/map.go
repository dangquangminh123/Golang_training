package main
import (
	"fmt"
)

func main() {
	/* map[string]float64
	 là một cấu trúc dữ liệu để lưu trữ nhiều món hàng và giá của chúng. Cụ thể:
		Khóa (key): Là tên món hàng (ví dụ: "soup", "pie").
		Giá trị (value): Là giá tiền của món hàng (ví dụ: 4.99, 6.99).
	*/
	menu := map[string]float64 {
		"soup": 4.99,
		"pie": 6.99,
		"salad": 8.00,
		"toffee pudding": 3.55,
	}

	// Lấy toàn bộ danh sách
	fmt.Println(menu)
	// Truy xuất dữ liệu qua key
	fmt.Println(menu["pie"])

	// Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		01422354: "PHP",
		40238127: "NodeJs",
		94802324: "NestJs",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[01422354])
 //Cập nhập lại giá trị mới cho key

	phonebook[40238127] = "Laravel"
	fmt.Println(phonebook)
 //Cập nhập lại giá trị mới cho key
	phonebook[94802324] = "Java"
	fmt.Println(phonebook)
}