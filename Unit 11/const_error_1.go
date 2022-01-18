package main

import "fmt"

func main() {
	const age int = 10
	const name string = "Maria"
	const score int // Compiler error

	age = 20       // Compiler error
	name = "Grace" // Compiler error
}
