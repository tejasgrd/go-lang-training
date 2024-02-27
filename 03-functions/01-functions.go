package main

import "fmt"

func main() {
	sayHi()
	greet("Tejas")
	greetUser("Tejas", "Garde")
	fmt.Print(getGreetMsg("Tejas", "Garde"))
	fmt.Print(divide(5, 2))
	fmt.Println()
	q, r := divide(5, 2)
	fmt.Printf("quotent = %d & remender = %d \n", q, r)
}

func sayHi() {
	fmt.Print("Hi There\n")
}

func greet(name string) {
	fmt.Printf("Hi %s, have a nice Day\n", name)
}

func greetUser(name string, last_name string) {
	fmt.Printf("Hi %s %s, have a nice Day\n", name, last_name)
}

func getGreetMsg(name string, last_name string) string {
	return fmt.Sprintf("Hi %s %s, have a nice Day\n", name, last_name)
}

// func divide(x, y int) (int, int) {
// 	q := x / y
// 	r := x % y
// 	return q, r
// }

// Using Named Result Sytax
func divide(x, y int) (q int, r int) {
	q = x / y
	r = x % y
	return
}
