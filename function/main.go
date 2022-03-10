package main

import "fmt"

func main() {
	fmt.Println("aaaa")
}

//===No Return Value===//
func show(fristname, lastname string) {
	fmt.Println("Hello", fristname, lastname)
}

//===Return One Value===//
func sum(a, b int) int {
	return a + b
}

//===Return Many Value===//
func swap(a, b int) (int, int) {
	return b, a
}
