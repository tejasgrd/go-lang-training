package main

import "fmt"

func main() {
	// const (
	// 	red   int = 0
	// 	green int = 1
	// 	blue  int = 2
	// )

	// const (
	// 	red = iota
	// 	green
	// 	blue
	// )

	const (
		red = iota * 2
		green
		blue
	)

	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
	)

	fmt.Printf("red = %d, green =% d , blue = %d", red, green, blue)
	fmt.Printf("%b %b %b %b %b ", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG)
	fmt.Printf("%d %d %d %d %d ", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG)
}
