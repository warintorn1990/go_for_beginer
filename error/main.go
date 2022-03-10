package main

import "errors"

func main() {
	var s = "s"

	if s == nil {
		return errors.New("s is nil")
	}
}
