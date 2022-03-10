package main

import "fmt"

func main() {
	var a []int
	a = append(a, 1, 2, 3, 4, 5)
	a = append(a, 6)
	a = append(a, 7)

	b := len(a)
	c := cap(a)

	fmt.Println(a)
	fmt.Println("Len", b)
	fmt.Println("Cap", c)
}
