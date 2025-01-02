package main

import (
	"fmt"
	// "strings"
)

func main() {
	var ages = [3]int{20, 25, 30}
	/* :=dùng để khai báo và khởi tạo biến mới trong cùng một lúc. 
	không cần chỉ định kiểu dữ liệu của biến vì Go tự động suy luận dựa trên giá trị được gán.
	*/
	/* [n]type (với n là số phần tử, type là kiểu dữ liệu). */
	names := [4]string{"php", "laravel", "nodejs", "reactjs"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use array under hood)
	/* Slice: Là cấu trúc dữ liệu sử dụng mảng nhưng có thể thay đổi kích thước động (khác với mảng có kích thước cố định).
	Cú pháp: []type (không cần định nghĩa số phần tử). */ 
	var scores = []int{100, 50, 60}
	scores[2] = 25

	/* append Công dụng: Thêm phần tử mới vào cuối slice và trả về một slice mới với kích thước tăng lên. 
	+Hàm append không thay đổi slice gốc, mà trả về một slice mới., 
	+Bạn có thể thêm nhiều phần tử cùng lúc*/ 
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// Slice ranges
	/*
		+slice[start:end]: Lấy các phần tử từ chỉ mục start đến end-1.
		+slice[start:]: Lấy các phần tử từ chỉ mục start đến hết slice.
		+slice[:end]: Lấy các phần tử từ đầu slice đến chỉ mục end-1.
	*/
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Nest");
	fmt.Println(rangeOne)
}