package main

import (
	"errors"
	"fmt"
)

// error variable starts with Err
var ErrFileOpen error = errors.New("error opening file")
var ErrorDivideByZero error = errors.New("divide by zero")

func main() {
	var err error
	// err = errors.New("dummy error")
	// if err != nil {
	// 	fmt.Println("Some error occered")
	// } else {
	// 	fmt.Println("all good")
	// }

	err = ErrFileOpen
	if err == ErrFileOpen {
		fmt.Println("Error opening file")
	} else if err == ErrorDivideByZero {
		fmt.Println("divide by zero")
	} else {
		fmt.Println("Unknown Error")
	}
}
