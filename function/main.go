package main

import "fmt"

func main() {

	show("Arm", "Art")

	s := sum(1, 1)
	fmt.Println(s)

	a, b := swap(5, 10)

	fmt.Println("a", a)
	fmt.Println("a", b)
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
