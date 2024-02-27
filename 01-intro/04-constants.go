package main

import "fmt"

func main() {
	const app_version string = "1.0"

	const (
		pi          = "3.14"
		new_version = "2.0"
	)

	fmt.Println(pi, new_version)

	fmt.Print(app_version)
}
