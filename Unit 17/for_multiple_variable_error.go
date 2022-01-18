package main

import "fmt"

func main() {
	// Vế sau đã định nghĩa sai 
	// Phải tách rởi ra để định nghĩa
	for i, j := 0, 0; i < 10; i++, j+=2 { // 컴파일 에러. syntax error: unexpected comma, expecting {
		fmt.Println(i, j)
	}
}
