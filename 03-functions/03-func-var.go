package main

import "fmt"

func main() {
	sayHi := func() {
		fmt.Println("Hi There from Anonymous")
	}

	sayHi()
	var quotent_temender func(int, int) (int, int)
	quotent_temender = func(x, y int) (q int, r int) {
		q = x / y
		r = x % y
		return
	}

	fmt.Print(quotent_temender(50, 10))

	//pass fucntions as argument to other functions

}
