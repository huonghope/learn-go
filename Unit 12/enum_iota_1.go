package main

import "fmt"

func main() {

	// Các giá trị phía sau sẽ gắn tự động tăng lên 1 đơn vị
	const (
		Sunday       = iota // 0
		Monday              // 1
		Tuesday             // 2
		Wednesday           // 3
		Thursday            // 4
		Friday              // 5
		Saturday            // 6
		numberOfDays        // 7
	)

	fmt.Println(Thursday)     // 4
	fmt.Println(numberOfDays) // 7
}
