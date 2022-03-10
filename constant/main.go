package main

import "fmt"

func main() {
	const value int = 5
	const title = "Go For Brginner"

	const (
		january = iota + 1
		february
		march
		april
		may
		june
		july
		august
		september
		october
		november
		december
	)

	fmt.Println("value", value)
	fmt.Println("title", title)

	fmt.Println("january ==>", january)
	fmt.Println("february ==>", february)
	fmt.Println("march ==>", march)
	fmt.Println("april ==>", april)
	fmt.Println("may ==>", may)
	fmt.Println("june ==>", june)
	fmt.Println("july ==>", july)
	fmt.Println("august ==>", august)
	fmt.Println("september ==>", september)
	fmt.Println("october ==>", october)
	fmt.Println("november ==>", november)
	fmt.Println("december ==>", december)
}
