package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

// Phải dùng pointer để sau khi tính giá trị thì ta save ngược lại vào biến area luôn
func rectangleArea(rect *Rectangle) int { // 매개변수로 구조체 포인터를 받음
	return rect.width * rect.height
}

func main() {
	rect := Rectangle{30, 30}

	// Áp dụng ưu điểm của pointers, để giá trị sau khi tính xong sẽ lưu ngược lại vào rectangleArea luôn
	area := rectangleArea(&rect) // 구조체의 포인터를 넘김

	fmt.Println(area) // 900
}
