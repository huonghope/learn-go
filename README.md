### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2007" target="_blank"> Unit 7: Variable </a>
```go
func main() {
	a, b, c, d := 1, 3.4, "Hello, world!", false // a = 1, b = 3.4, c = "Hello, world!", d = false

	fmt.Println(a, b, c, d)
}
```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2008" target="_blank"> Unit 8: Byte, Complex, Float-rouding, float, int (max-min), number_conversion, rune </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2009" target="_blank"> Unit 9: String, String length, String Operation </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2010" target="_blank"> Unit 10: Boolean logical, boolean operator </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2011" target="_blank"> Unit 11: Const value, constan error </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2012" target="_blank"> Unit 12: Enum, Enum-automatic, Enum-increment </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2013" target="_blank"> Unit 13: Simple Opeator, Operator Priority </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2014" target="_blank"> Unit 14: Package import, No-use package import  </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2015" target="_blank"> Unit 15: If-else, If-error </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2016" target="_blank"> Unit 16: If-func scope </a>
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
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2017" target="_blank"> Unit 17: For loop, for-continue, for-break, for-error, for-infinitie </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2018" target="_blank"> Unit 18: Goto </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2019" target="_blank"> Unit 19: Switch, switch-expression, swith-multiple </a>
```go
package main

import "fmt"

func main() {
	i := 3

	switch i {
	case 2, 4, 6: // i가 2, 4, 6일 때
		fmt.Println("짝수")
	case 1, 3, 5: // i가 1, 3, 5일 때
		fmt.Println("홀수")
	}

  for i := 1; i <= 100; i++ {
		switch {                         // case에 조건식을 지정했으므로 판단할 변수는 생략
		case i%3 == 0 && i%5 == 0:       // 3의 배수이면서 5의 배수일 때
			fmt.Println("FizzBuzz")  // FizzBuzz 출력
		case i%3 == 0:                   // 3의 배수일 때
			fmt.Println("Fizz")      // Fizz 출력
		case i%5 == 0:                   // 5의 배수일 때
			fmt.Println("Buzz")      // Buzz 출력
		default:                         // 아무 조건에도 해당하지 않을 때
			fmt.Println(i)           // 숫자 출력
		}
	}
}
```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2020" target="_blank"> Unit 20: For-switch,case-multi </a>
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2020" target="_blank"> Unit 21: Array, Array for len, Array for range </a>
```go
package main

import "fmt"

func main() {
	var myArray [5]int // int형이며 길이가 5인 배열 선언

	myArray[2] = 7     // 배열의 세 번째 요소에 7 대입

  var a [5]int = [5]int{32, 29, 78, 16, 81} // int형이며 길이가 5인 배열을 선언하고 초기화
	var b = [5]int{32, 29, 78, 16, 81}        // 배열을 선언할 때 자료형과 길이 생략
	c := [5]int{32, 29, 78, 16, 81}           // 배열을 선언할 때 var 키워드, 자료형과 길이 생략


	bArray := [5]int{
		32,
		29,
		78,
		16,
		81, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}

	cArray := [...]int{32, 29, 78, 16, 81} // 초기화할 요소가 5개이며 ...을 사용했으므로 
	                                  // 배열 크기는 5로 설정됨

	dArray := [...]string{"Maria", "Andrew", "Jonh"} // 초기화할 요소가 3개이며 ...을 사용했으므로 
	                                            // 배열 크기는 3으로 설정됨

	fmt.Println(myArray) // [0 0 7 0 0]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(aArray)
	fmt.Println(bArray)
	fmt.Println(cArray)


  a := [5]int{32, 29, 78, 16, 81}

	for value := range a { // value에는 값 대신 인덱스가 들어감
		fmt.Println(value)
	}

  for _, value := range a { // 인덱스는 생략, value에 배열 요소의 값이 들어감
		fmt.Println(value)
	}

  for i, value := range a { // i에는 인덱스, value에는 배열 요소의 값이 들어감
		fmt.Println(i, value)
	}

  a := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(a)


}

```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2022" target="_blank"> Unit 19: Slice, Slice copy, len, Slice-sclicing </a>
```go
package main

import "fmt"

func main() {
	var a []int = make([]int, 10) // make 함수로 int형에 길이가 5인 슬라이스에 공간 할당
	var b = make([]int, 5)       // 슬라이스를 선언할 때 자료형과 [] 생략
	c := make([]int, 5)          // 슬라이스를 선언할 때 var 키워드, 자료형과 [] 생략

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	a = append(a, b...) // 슬라이스 a에 슬라이스 b를 붙일 때는 b...을 씀

	fmt.Println(a) // [1 2 3 4 5 6]

	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a[:])        // [1 2 3 4 5]
	fmt.Println(a[0:])       // [1 2 3 4 5]
	fmt.Println(a[:5])       // [1 2 3 4 5]
	fmt.Println(a[0:len(a)]) // [1 2 3 4 5]

	fmt.Println(a[3:])       // [4 5]
	fmt.Println(a[:3])       // [1 2 3]
	fmt.Println(a[1:3])      // [2 3]


// Slice Slicing Cap example
	b := a[0:6:8] // 인덱스 0부터 6까지 가져와서 부분 슬라이스로 만들고 용량을 8로 설정

	fmt.Println(len(b), cap(b)) // 6 8: 길이가 6이며 용량이 8인 슬라이스
	fmt.Println(b)              // [1 2 3 4 5 6]
}
```