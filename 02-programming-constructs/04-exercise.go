package main

import (
	"fmt"
)

func main() {
LOOP:
	for i := 2; i <= 100; i++ {

		for j := 2; j <= (i / 2); j++ {
			if i%j == 0 {
				continue LOOP
			}
		}

		fmt.Printf(" %d is Prime Number \n", i)

	}
}
