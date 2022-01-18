package main

import "fmt"

func main() {

	// các giá trị sau sẽ bằng giá trị iota tăng lên 1 đơn vị và nhân lên với 30
	const (
		a = iota * 30 // a == 0 (0 * 30)
		b = iota * 30 // b == 30 (1 * 30)
		c = iota * 30 // c == 60 (2 * 30)
		d = iota * 30 // d == 90 (3 * 30)
	)

	fmt.Println(a, b, c, d)
}
