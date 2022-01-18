package main

import "fmt"

func main() {

	// Định nghĩa trước các kiểu của biến số, Ở đây là kiểu INT
	var x, y int
	var age int

	x, y, age = 10, 20, 5 // x = 10, y = 20, age = 5: 병렬 할당

	fmt.Println(x, y, age)
}
