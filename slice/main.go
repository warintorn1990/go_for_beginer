package main

import "fmt"

func main() {
	var num []int
	num = make([]int, 4)

	num[0] = 0
	num[1] = 1
	num[2] = 2

	fmt.Println(num)
}
