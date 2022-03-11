package main

import "fmt"

func main() {
	// a := make(map[string]string)
	// a := map[string]string{}

	// a["a"] = "Naruto"

	a := map[string]string{
		"a": "Naruto",
		"b": "Onepice",
		"c": "DragonBall",
	}

	fmt.Println(a["a"])
	fmt.Println(a["b"])
	fmt.Println(a["c"])
}
