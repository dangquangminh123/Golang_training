package main

import "fmt"

// func updateName(n string) {
// 	n = "wedge"
// }

// con trỏ đến 1 giá tri chuỗi trong hàm
// *string là giá trị con trỏ trỏ đến là một chuỗi (string).
func updateName(n *string) {
	// hay đổi giá trị tại địa chỉ bộ nhớ của name
	*n = "PHP"
}

func main() {
	name := "Laravel"

	// updateName(name)
	// fmt.Println(name)
	
	// &name: Lấy địa chỉ bộ nhớ nơi lưu trữ giá trị của biến name.
	fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name
	// m: In ra địa chỉ bộ nhớ mà con trỏ chỉ đến (giống &name).
	fmt.Println("memory address:", m)

	// *m: Lấy giá trị tại bộ nhớ mà m trỏ đến (là giá trị của name).
	fmt.Println("value at memory address:", *m)

	updateName(m) // Thay vì truyền trực tiếp giá trị, bạn truyền địa chỉ bộ nhớ của name (m).
	fmt.Println(name)

}