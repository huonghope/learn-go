package main

import "fmt"

// Một struct được định nghĩa có 2 giá trị chính
type Rectangle struct {
	width  int
	height int
}

// Biến đằng trước tên là 1 class, thì hàm được hiểu là 1 method, và con trò sẽ đóng vai trò là this để thay đổi giá của giá biến struct đó
// Các thay đổi trong hàm area sẽ được phản ánh qua biến khai báo tạo hàm Main
//          ↓ 리시버 변수 정의(연결할 구조체 지정)
func (rect *Rectangle) area() int {
	return rect.width * rect.height
	//     ↑  리시버 변수를 사용하여 현재 인스턴스에 접근할 수 있음
}

func main() {
	rect := Rectangle{10, 20}

	fmt.Println(rect.area()) // 200
}
