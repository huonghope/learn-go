package main

import "fmt"

func main() {
	var s1 string = "Hello"
	var s2 string = "World"
	var s3 string = "Go"

	fmt.Println(s1 == s2)          // false
	fmt.Println(s1 + s3 + s2)      // Hello World Go
	fmt.Println("I said" + s1) // I said Hello
}
