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
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2021" target="_blank"> Unit 21: Array, Array for len, Array for range </a>
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
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2022" target="_blank"> Unit 22: Slice, Slice copy, len, Slice-sclicing </a>
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
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2023" target="_blank"> Unit 23: Map, Map-delete, Map-range, Map-in-map </a>
```go
package main

import "fmt"

func main() {

	// Dùng make để tạo map
	var a map[string]int = make(map[string]int) // make 함수로 키는 string 값은 int인 맵에 공간 할당
	var b = make(map[string]int)                // 맵을 선언할 때 map 키워드와 자료형 생략
	c := make(map[string]int)                   // 맵을 선언할 때 var, map 키워드와 자료형 생략

	fmt.Println(a) //map[]
	fmt.Println(b) //map[]
	fmt.Println(c) //map[]

	// Khai báo map với string là keyword và giá trị là int
	a := map[string]int{"Hello": 10, "world": 20}

	b := map[string]int{
		"Hello": 10,
		"world": 20, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}

	// Dùng make để khai báo và tạo map, Sau đó sẽ khai báo các giá trị cơ bản
	solarSystem := make(map[string]float32) // 키는 string, 값은 float32인 맵 생성 및 공간 할당
	solarSystem["Mercury"] = 87.969 // 맵[키] = 값
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528

	fmt.Println(solarSystem["Earth"]) // 365.25641

	// Nếu k muốn sử dụng key thì dùng toán tử _
	for _, value := range solarSystem { // 키 변수를 사용하고 싶지 않다면 _ 사용
		fmt.Println(value)
	}

	// ok sẽ trả về 2 giá trị là true và false, nếu là true có nghĩ là với giá trị key đó sẽ trả về 1 giá trị tương đương
	// ngược lại nếu là false thì sẽ không có giá trị tương đương
	value, ok := solarSystem["Pluto"] // 맵에 키가 있는지 검사할 때는 리턴값을 두 개 활용
	fmt.Println(value, ok)            // 0 false: 맵에 키가 없으면 두 번째 리턴값으로 false가 리턴됨

	if value, ok := solarSystem["Saturn"]; ok {
		fmt.Println(value) // 10756.1995
	}

	// Dùng toán tử delete để xóa 1 biến trong map, với giá trị keyword truyền vào
	a := map[string]int{"Hello": 10, "world": 20}
	delete(a, "world") // 맵 a에서 world 키 삭제
	

	// Khai báo map in map 
	terrestrialPlanet := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022E+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.8676E+24,
			"orbitalPeriod": 224.70069,
		},
		"Earth": map[string]float32{
			"meanRadius":    6371.0,
			"mass":          5.97219E+24,
			"orbitalPeriod": 365.25641,
		},
		"Mars": map[string]float32{
			"meanRadius":    3389.5,
			"mass":          6.4185E+23,
			"orbitalPeriod": 686.9600,
		},
	}

	fmt.Println(terrestrialPlanet["Mars"]["mass"]) // 6.4185E+2
}

```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2024" target="_blank"> Unit 24: Function, function-map </a>
```go
package main

import "fmt"

func main() {
	func() { // 함수에 이름이 없음
		fmt.Println("Hello, world!")
	}()

	func(s string) {   // 익명 함수를 정의한 뒤
		fmt.Println(s)
	}("Hello, world!") // 바로 호출

	// Định nghĩa hàm và gọi hàm luôn, truyền vào 2 tham số tương đương là 1 và 2
	r := func(a int, b int) int { // 익명 함수를 정의한 뒤
		return a + b
	}(1, 2)                       // 바로 호출하여 리턴값을 변수 r에 저장

	fmt.Println(r) // 3
}

// Thêm kiểu dữ liệu ở sau giấy () để biểu trị kiểu của biến return
func sum(a int, b int) int { // int형 매개변수 a, b 그리고 int 형 리턴값을 가지는 함수 정의
	return a + b
}
```
```go
package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	// Khai báo biến f, gồm 2 key là sum và diff, value của 2 key này chính là phần định nghĩa của 2 hàm
	// khi gọi hàm sẽ truyền vào 2 giá trị để nhận lại kết quả của 2 hàm
	f := map[string]func(int, int) int{ // 함수를 저장할 수 있는 맵을 생성한 뒤 함수로 초기화
		"sum":  sum,
		"diff": diff,
	}

	fmt.Println(f["sum"](1, 2))  // 3: 맵에 sum 키를 지정하여 함수 호출
	fmt.Println(f["diff"](1, 2)) // -1: 맵에 diff 키를 지정하여 함수 호출
}
```
```go
package main

import "fmt"

// Định nghĩa 2 hàm số, nhưng sẽ trả về 2 giá trị
func SumAndDiff(a int, b int) (int, int) { // int형 리턴값이 두 개인 함수 정의
	return a + b, a - b // 리턴할 값 또는 변수를 차례대로 나열. 합과 차를 동시에 리턴
}

// Định nghĩa trước 2 biến return, sẽ tự động return ra 2 biến này trong quá trình tính logic trong hàm
func SumAndDiff(a int, b int) (sum int, diff int) { // 리턴값을 각각 sum, diff로 이름을 정함
	sum = a + b  // 리턴값 변수 sum에 합 대입
	diff = a - b // 리턴값 변수 diff에 차 대입
	return
}

func hello() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5 // 리턴할 값을 차례대로 나열
}

// Cách gọi hàm đồng quy
func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	sum, diff := SumAndDiff(6, 2) // 변수 두 개에 리턴값 두 개를 대입
	fmt.Println(sum, diff)        // 8 4: 합과 차

	// Với trường hợp nếu không muốn sử dụng biến đầu tiên, dùng toán tử _ để k quan tâm
	_, diff := SumAndDiff(6, 2) // 첫 번째 리턴값은 _으로 생략, 두 번째 리턴값만 사용
	fmt.Println(diff)           // 4: 차

	// Tương tự như ví dụ trên
	a, _, c, _, e := hello() // 2, 4번째 리턴값은 생략
	fmt.Println(a, c, e)     // 1 3 5
}
```
```go
package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	// Dùng slice để định nghĩa 1 array, gồm 2 mảng
	// Giá trị thứ 0 của mảng là hàm sum
	// Giá trị thứ 1 của mảng là hàm diff
	// Khi gọi đến giá trị, truyền vào biến sẽ nhận lại được giá trị tính toán của hàm
	f := []func(int, int) int{sum, diff} // 함수를 저장할 수 있는 슬라이스를 생성한 뒤 함수로 초기화

	fmt.Println(f[0](1, 2)) // 3: 배열의 첫 번째 요소로 함수 호출
	fmt.Println(f[1](1, 2)) // -1: 배열의 두 번째 요소로 함수 호출
}
```
```go
package main

import "fmt"

// n sẽ biểu thị nhận vào tất cả các giá trị khi truyền vào qua hàm
// giống như js6
func sum(n ...int) int { // int형 가변인자를 받는 함수 정의
	total := 0
	for _, value := range n { // range로 가변인자의 모든 값을 꺼냄
		total += value   // 꺼낸 값을 모두 더함
	}

	return total
}

func sumNumber(a int, b int) int {
	return a + b
}

func main() {
	r := sum(1, 2, 3, 4, 5)
	fmt.Println(r) // 15

	// Lưu ý dấu ... ở đằng sau tên biến
	n := []int{1, 2, 3, 4, 5}
	r := sum(n...) // ...를 사용하여 가변인자에  // 슬라이스를 바로 넘겨줌
	fmt.Println(r) // 15

	// Khai báo biến hàm
	// int được hiểu là biến sẽ lưu kết quả của biếm số và sumNumber là hàm sẽ triển khai thaanc của hàm khai báo trước đó
	var hello func(a int, b int) int = sumNumber // 함수를 저장하는 변수를 선언하고 함수 대입
	world := sumNumber                           // 변수 선언과 동시에 함수를 바로 대입

	fmt.Println(hello(1, 2)) // 3: hello 변수에 저장된 sum 함수 호출
	fmt.Println(world(1, 2)) // 3: world 변수에 저장된 sum 함수 호출
}

```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2025" target="_blank"> Unit 25: Closure </a>
```go
package main
import "fmt"

// Hàm trả về là 1 hàm số không tên, với tiền return giá trị int
//              ↓ 리턴 값이 익명 함수
func calc() func(x int) int {
	a, b := 3, 5           // 지역 변수는 함수가 끝나면 소멸되지만
	return func(x int) int {
		return a*x + b // 클로저이므로 함수를 호출 할 때마다 변수 a와 b의 값을 사용할 수 있음
	}
	// ↑ 익명 함수를 리턴
}


//    ↓ 함수 안에서
func main() {
	//      ↓ 익명 함수를 선언 및 정의
	sum := func(a, b int) int {
		return a + b
	}

	r := sum(1, 2) // 익명함수 사용

	fmt.Println(r) // 3

	// Định nghĩa hàm và nhận 2 giá trị ở bên ngoài để tính toán luôn
	a, b := 3, 5
	f := func(x int) int {
		return a*x + b // 함수 바깥의 변수 a, b 사용
	}

	y := f(5)
	fmt.Println(y) // 20

	f := calc() // calc 함수를 실행하여 리턴값으로 나온 클로저를 변수에 저장
	fmt.Println(f(1)) // 8
	fmt.Println(f(2)) // 11
	fmt.Println(f(3)) // 14
	fmt.Println(f(4)) // 17
	fmt.Println(f(5)) // 20
}

```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2026" target="_blank"> Unit 26: Defer </a>
```go
package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}
func HelloWorld() {
	//↓ HelloWorld() 함수가 끝나기 직전에 호출
	defer func() {
		fmt.Println("world")
	}()

	func() {
		fmt.Println("Hello")
	}()
}

func ReadHello() {
	// Dùng defer để đóng file đã đọc ở cuối hàm
	file, err := os.Open("hello.txt")
	defer file.Close() // 지연 호출한 file.Close()가 맨 마지막에 호출됨

	if err != nil {
		fmt.Println(err)
		return // file.Close() 호출
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return // file.Close() 호출
	}

	fmt.Println(string(buf))
	// file.Close() 호출
}

func main() {
        //↓ 함수의 호출이 지연됨.
	defer world() // 현재 함수(main())가 끝나기 직전에 호출
	hello()
	hello()
	hello()
	HelloWorld()

	// 4 3 2 1 0 : sẽ bị đảo ngược
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	ReadHello()
}
```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2027" target="_blank"> Unit 27: Panic, Panic-recover </a> 
```go
package main

import "fmt"

func f() {
	defer func() {         // recover 함수는 지연 호출로 사용해야 함
		s := recover() // 패닉이 발생해도 프로그램을 종료하지 않음, panic 함수에서 설정한 에러 메시지를 받아옴
		fmt.Println(s)
	}()

	panic("Error !!!") // panic 함수로 에러 메시지 설정, 패닉 발생
}
func checkingRecover() {
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
	checkingRecover()
	fmt.Println("Hello, world!") // 패닉이 발생했지만 계속 실행됨

	panic("Error !!")
	fmt.Println("Panic error") // 실행되지 않음
}
```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2028" target="_blank"> Unit 28: Pointer </a> 
```go
package main

import "fmt"

func hello(n *int) {
	*n = 2 // 포인터 변수 n를 역참조하여 메모리에 2를 대입
}

func main() {
	var numPtr *int // 포인터형 변수를 선언하면 nil로 초기화됨
	var numPtrInit *int = new(int)

	fmt.Println(numPtr) // nil
	fmt.Println(numPtrInit) // 0xc0820062d0: 메모리 주소. 시스템 마다, 실행할 때마다 달라짐

	var numPtrInt *int = new(int)
	*numPtrInt = 1
	fmt.Println(*numPtrInt) //1: 포인터 형 변수에서 값을 가져오기

	// Vì ta truyền giá trị tại vị trí memory của biến n, lên biến n ra khỏi hàm cũng sẽ bị thay đổi giá trị
	var n int = 1
	hello(&n) // 1이 들어있는 변수 n의 메모리 주소를 hello 함수에 넘김
	fmt.Println(n) // 2: hello 함수에서 n의 메모리 공간에 2를 대입했으므로 바깥에 있는 n의 값이 바뀌었음

	var num int = 1
	var numPtr *int = &num //참조로 num변수의 메모리 주소를 구하여 - numPtr 포인터 변수에 대입함
	
	fmt.Println(numPtr) // 0xc0820062d0: numPtr 포인터 변수에 저장된 메모리 주소
	fmt.Println(&num) // 0xc0820062d0: 참조로 num 변수의 메모리 주소를 구함
}
```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2029" target="_blank"> Unit 29: Struct, struct-constructor, struct-func-argument </a>
```go
package main

import "fmt"

// Khai báo kiểu dữ liệu struct, c/c++ same
// Gồm 2 biến giá trị là w & h
type Rectangle struct {
	width  int
	height int
}

// Khai báo 1 hàm nhận vào 2 giá trị w và h, Sau đó tạo ra 1 biến theo kiểu struct và trờ về giá trị địa chỉ
// Ở hàm main sẽ lưu địa chỉ đó vào, nhưng là 1 biến mới được tạo thành
func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height} // 구조체 인스턴스를 생성한 뒤 포인터를 리턴
}

func rectangleScaleA(rect *Rectangle, factor int) { // 매개변수로 구조체 포인터를 받음
	rect.width = rect.width * factor   // 포인터이므로 원래의 값이 변경됨
	rect.height = rect.height * factor // 포인터이므로 원래의 값이 변경됨
}

func main() {
	// Khai báo biến kiểu pointer
	var rect1 *Rectangle   // 구조체 포인터 선언
	rect1 = new(Rectangle) // 구조체 포인터에 메모리 할당

	rect2 := new(Rectangle) // 구조체 포인터 선언과 동시에 메모리 할당

	fmt.Println(rect1)
	fmt.Println(rect2)

	// Các cách khai báo giá trị cho kiểu struct
	var rect1 Rectangle = Rectangle{10, 20} // 구조체 인스턴스를 생성하면서 값 초기화

	rect2 := Rectangle{45, 62} // var 키워드 생략 구조체 인스턴스를 생성하면서 값 초기화

	// Truyền 2 biến giá trị vào theo kiểu tường minh
	rect3 := Rectangle{width: 30, height: 15} // 구조체 필드를 지정하여 값 초기화

	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)

	// 2 kiểu khai báo là giống nhau 
	// Thông qua hàm khởi tạo ta đã định nghĩa
	rect := NewRectangle(20, 10)
	fmt.Println(rect) // &{20 10}

	// Trực tiếp thông qua tử &
	rect := &Rectangle{20, 10} // 구조체를 초기화한 뒤 메모리 주소를 대입
	fmt.Println(rect) // &{20 10}

	rectangleScaleA(&rect1, 10) // 구조체의 포인터를 넘김
	fmt.Println(rect1)          // {300 300}: rect1에 바뀐 값이 들어감
}
```
### <a href="https://github.com/huonghope/learn-go/tree/master/Unit%2030" target="_blank"> Unit 30: Struct, struct-method</a>
```go

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

// Toán tử _ với ý đồ ngầm định là không sử dụng
func (_ Rectangle) dummy() { // _로 리시버 변수 생략
	fmt.Println("dummy")
}

func main() {
	rect := Rectangle{10, 20}

	fmt.Println(rect.area()) // 200

	// Chỉ in ra 1 hàm print đơn giản
	var r Rectangle
	r.dummy() // dummy
}


```