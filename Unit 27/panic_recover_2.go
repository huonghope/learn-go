package main

import "fmt"

func f() {
	// Ở phần dưới có error tuy nhiên thông qua recover() đã khắc phục và chương trình vẫn chạy bthg
	defer func() {
		s := recover()   // recover 함수로 런타임 에러(패닉) 상황을 복구
		fmt.Println(s)
	}()

	a := [...]int{1, 2}      // 정수가 2개 저장된 배열

	//  biến a chỉ có 2 phần từ những ở đây đã xuất đến 5 phần tử
	for i := 0; i < 5; i++ { // 배열 크기를 벗어난 접근 ERROR
		fmt.Println(a[i])
	}
}

func main() {
	f()

	fmt.Println("Hello, world!") // 런타임 에러가 발생했지만 recover 함수로 복구되었기 때문에 이 부분은 정상적으로 실행됨
}
