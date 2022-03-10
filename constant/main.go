package main

import "fmt"

func main() {
	const (
		sunday = iota
		monday
		tuesday
	)
	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(tuesday)
}
