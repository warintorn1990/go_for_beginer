package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	variadicString(s...)

	//variadicString("a", "b", "c")
}

func variadicString(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
