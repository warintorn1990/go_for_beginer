package main

import "fmt"

func main() {

	//===Loop With 3 Component
	for i := 0; i < 10; i++ {
		fmt.Println("i ==> ", i)
	}

	//===Loop With Condition===//
	var k = 0
	for k < 10 {
		fmt.Println("k ==> ", k)
		k++
	}

	//===Infinity Loop===//
	for {
		fmt.Println("===Infinity===")
	}
}
