package main

import (
	"fmt"
	"strconv"
)

func main() {

	//===IF ELSE===//
	name := "arm"
	if name != "" {
		fmt.Println("Hi", name)
	} else {
		fmt.Println("Hi Every One")
	}

	//===IF WITHOUT SHORT STATEMENR===//
	n, err := strconv.Atoi("5s")

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(n)
	}

	//===IF WITH SHORT STATEMENT===//
	if n, err := strconv.Atoi("5s"); err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println(n)
	}

}
