### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2007" target="_blank"> Unit 7: Variable
```go
func main() {
	a, b, c, d := 1, 3.4, "Hello, world!", false // a = 1, b = 3.4, c = "Hello, world!", d = false

	fmt.Println(a, b, c, d)
}
```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2008" target="_blank"> Unit 8: Byte, Complex, Float-rouding, float, int (max-min), number_conversion, rune
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2009" target="_blank"> Unit 9: String, String length, String Operation
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2010" target="_blank"> Unit 10: Boolean logical, boolean operator
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2011" target="_blank"> Unit 11: Const value, constan error
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2012" target="_blank"> Unit 12: Enum, Enum-automatic, Enum-increment
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2013" target="_blank"> Unit 13: Simple Opeator, Operator Priority
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2014" target="_blank"> Unit 14: Package import, No-use package import 
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2015" target="_blank"> Unit 15: If-else, If-error
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2016" target="_blank"> Unit 16: If-func scope
```go
package main

import "fmt"

func main() {
  // Đọc từ file ReadFile, và đồng thời xét xem err có lỗi không
  // Nếu error không có lỗi thì đưa vào if loop đầu tiên
  // Tuy nhiên có 1 lưu ý là những biến định nghĩa trên lệnh if này, chỉ sử dụng được ở trong mệnh đề if, nếu sử dụng ngoài sẽ phát sinh lỗi
	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b) // 변수 b를 사용할 수 있음
	} else {
		fmt.Println(err) // 변수 err을 사용할 수 있음
	}

	fmt.Println(b)   // 변수 b를 사용할 수 없음. 컴파일 에러
	fmt.Println(err) // 변수 err을 사용할 수 없음. 컴파일 에러

  var b []byte
	var err error

  // Để giải quyết vấn đề trên tả có thể định nghĩa 2 biến đó và nhận từ thân lệnh if
	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b)
	}
}

```