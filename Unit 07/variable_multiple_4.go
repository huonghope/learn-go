package main

import "fmt"

func main() {

	// Định nghĩa 2 biến x,y kiểu int và age, name sẽ tự động định nghĩa kiểu
	var (
		x, y      int = 30, 50      // x와 y는 int형으로 결정
		age, name     = 10, "Maria" // age는 int, name은 string으로 결정
	)

	fmt.Println(x, y, age, name)
}
