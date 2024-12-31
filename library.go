package main

import (
	"fmt"
	"strings"
	"sort"
)

function main() {
	greeting := "Hello there varriant";

	// strings

	/*strings.Contains Kiểm tra xem chuỗi greeting có chứa chuỗi không ?
	strings.Contains(s, substr)
	+s: Chuỗi cần kiểm tra.
	+substr: Chuỗi con cần tìm.
	*/
	fmt.Println(strings.Contains(greeting, "Hello!"))
	
	/*strings.ReplaceAll Thay thế tất cả các xuất hiện của "Hello" trong chuỗi greeting bằng "hi". 
	strings.ReplaceAll(s, old, new)
	+s: Chuỗi cần thay thế.
	+old: Chuỗi cũ cần thay.
	+new: Chuỗi mới thay vào.
	*/
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi"))

	/* strings.ToUpper Công dụng: Chuyển toàn bộ chuỗi greeting thành chữ hoa.
	strings.ToUpper(s)
	+s: Chuỗi cần chuyển đổi.*/
	fmt.Println(strings.ToUpper(greeting))

	/* strings.Index Tìm vị trí xuất hiện đầu tiên của chuỗi con "th" trong chuỗi greeting.
	strings.Index(s, substr)
	+s: Chuỗi cần tìm. 
	+substr: Chuỗi con cần xác định vị trí.
	*/
	fmt.Println(strings.Index(greeting, "th"))

	/* strings.Split Tách chuỗi greeting thành mảng các chuỗi con, dựa trên ký tự " " (dấu cách).
	 strings.Split(s, sep)
	 +s: Chuỗi cần tách.
	 +sep: Ký tự phân tách.
	*/
	fmt.Println(strings.Split(greeting, " "))



	// package Sort
	// The original value is unchanted
	fmt.Println("original string value =", greeting)

	/* sort.Ints Sắp xếp slice ages (kiểu []int) theo thứ tự tăng dần.
	sort.Ints(a)
	+a: Slice số nguyên cần sắp xếp.
	*/
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)
	 
	/* sort.SearchInts Tìm chỉ mục của phần tử có giá trị 30 trong slice ages đã được sắp xếp.
	sort.SearchInts(a, x)
	+a: Slice đã được sắp xếp tăng dần
	+x: Giá trị cần tìm.
	*/ 
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"php", "laravel", "nodejs", "nest", "golang"}

	/*sort.Strings Sắp xếp slice names (kiểu []string) theo thứ tự tăng dần (alphabet).
	+sort.Strings(a)
	+a: Slice chuỗi cần sắp xếp.
	*/ 
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "nest"));


}