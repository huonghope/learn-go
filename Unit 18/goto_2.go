package main

import "fmt"

func main() {
	var a int = 1

	if a == 1 {
		goto ERROR // 여기까지만 실행하고 ERROR 레이블로 이동
		b := 1
		fmt.Println(b)
	}

// Sẽ chạy ngay khi if loop được đi vào
// Các câu lệnh bên dưới goto sẽ không được chạy nữa
ERROR:
	fmt.Println("Error")
}
