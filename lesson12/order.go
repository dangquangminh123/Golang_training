package main

import (
	"fmt"
	"os"
)

/*
	structs trong Go, một kiểu dữ liệu được định nghĩa,
 	cho phép nhóm các trường dữ liệu liên quan thành một kiểu duy nhất
*/ 
type bill struct {
	/* Struct là một tập hợp các trường (fields), mỗi trường có một tên và kiểu dữ liệu. */
	name string //Tên món
	items map[string]float64 //Lưu các món hàng, giá trị tiền bạc của chúng
	tip float64 //Tiền bo 
}

// Hàm trả về một giá trị kiểu bill
func newBill(name string) bill {
	b := bill {
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b;
}

func (b bill) addItem(name string, price float64) {
	/* b.items[name] = price chỉ cập nhật bản sao (copy) của mybill.items, không ảnh hưởng đến mybill gốc */
	b.items[name] = price
}

func (b *bill) format() string {
	fs := "Bill breakdown: \n";
	var total float64 = 0

	// List item
	for k, v := range b.items {
		/* 
		%-25v: In tên món hàng, căn trái với chiều dài tối đa là 25 ký tự.
		..$%v: In giá tiền món hàng.
		*/
		fs += fmt.Sprintf("%-25v ..$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ... $%0.2f", "tip:", b.tip)
	// total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total+b.tip)

	return fs
}

/* Pointer receiver (*bill) được dùng để thay đổi trực tiếp dữ liệu trong struct hoặc tránh sao chép dữ liệu lớn. */
/*Hàm này nhận con trỏ của mybill. 
	Truy cập b.tip thay đổi trực tiếp giá trị tip trong mybill.
	+b *bill là pointer receiver, thể hiện rằng hàm updateTip nhận một con trỏ đến kiểu dữ liệu bill làm receiver.
	+Dấu * trước tên kiểu dữ liệu (bill) chính là biểu hiện cho con trỏ.
*/
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}



//Save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".pdf", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}
/* Value receiver (bill) chỉ nên dùng khi:
Không cần thay đổi dữ liệu gốc.
Struct có kích thước nhỏ.*/