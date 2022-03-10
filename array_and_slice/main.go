package main

import "fmt"

func main() {

	var sixNum [6]int
	var nums []int

	//===1===//
	nums = sixNum[0:6]

	//===2===//
	// nums = sixNum[0:4]
	// nums = sixNum[:4]

	a := len(nums)
	b := cap(nums)

	fmt.Println(nums)
	fmt.Println(a)
	fmt.Println(b)

}
