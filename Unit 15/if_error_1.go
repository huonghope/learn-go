package main

import "fmt"

func main() {
	i := 10

	if i >= 5 // 컴파일 에러
	{
		fmt.Println("5 이상")
	}
	if i >= 5	{
		fmt.Println("5 이상")
	}
}