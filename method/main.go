package main

func main() {

	type Dog struct {
		name string
		age  int
	}

	// func Walk (d Dog) string{
	// 	return 'Walk...'
	// }

	func (d Dog) Walk(); string{
		pp := "Walk..."
		return pp
	}

}
