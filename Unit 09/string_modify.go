package main

import "fmt"

func main() {
	var s1 string = "Hello, world!\n"

	s1 = "abcdefg" // gắn giá trị cho biến s1

	fmt.Println(s1[0]) // print = 97

	s1[0] = 'z' // Thay đổi biến đầu của s1
}
