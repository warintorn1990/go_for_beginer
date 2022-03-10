package main

import "fmt"

type dog struct {
	name string
	age  int
}

func main() {

	// var lufy dog = dog{
	// 	name: "lufy",
	// 	age:  19,
	// }

	// var solo dog{
	// 	name: "Solo",
	// 	age:  22,
	// }

	sanji := dog{
		name: "Sanji",
		age:  22,
	}

	fmt.Println(sanji.name)

}
