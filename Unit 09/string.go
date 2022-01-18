package main

import "fmt"

func main() {
	var s1 string = "Hello, world!\n"
	var s2 string = "Hello go\n"
	var s3 string = "\ud55c\uae00"   //한글          
	var s4 string = "\U0000d55c\U0000ae00"  //한글    
	var s5 string = "\xed\x95\x9c\xea\xb8\x80" //한글

	var s7 string = `I said Hello, world!`

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s7)
}
