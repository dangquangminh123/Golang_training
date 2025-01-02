package main
import (
	"fmt"
)

func updateName(x string) string{
	x = "Developer"
	return x;
}

func updateMenu(y map[string]float64){
	y["coffee"] = 2.99
}

func main() {
	name := "NodeJs"

	name = updateName(name)
	fmt.Println(name)

	menu := map[string]float64 {
		"pie": 4.99,
		"ice cream": 2.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}