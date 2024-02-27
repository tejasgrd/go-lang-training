package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi There from Anonymous")
	}()

	func(name string) {
		fmt.Printf("Hi %s, have a nice Day\n", name)
	}("Tejas")

	msg := func(name string, last_name string) string {
		return fmt.Sprintf("Hi %s %s, have a nice Day\n", name, last_name)
	}("Tejas", "Garde")

	fmt.Printf("Message is -> %s", msg)
}
