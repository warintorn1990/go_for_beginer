package main

import "fmt"

func main() {

	var sixNum [6]int
	var nums []int

	//===1===//
	nums = sixNum[1:4]

	//===2===//
	nums = sixNum[0:4]
	nums = sixNum[:4]

	fmt.Println(nums)

}
