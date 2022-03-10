package main

import (
	"errors"
	"fmt"
)

func handleError(s string) error {
	var err error
	if s == "" {
		err = errors.New("s is nil")
	}

	return err
}

func main() {
	var s = ""

	err := handleError(s)

	fmt.Println(err)
}
