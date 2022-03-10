package main

import "errors"

func main() {
	var s = "s"

	if s == "" {
		return errors.New("s is nil")
	}
}
