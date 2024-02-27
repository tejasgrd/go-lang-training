package main

import "fmt"

func main() {
	//var nos [5]int
	//var nos [5]int = [5]int{1, 2, 3, 4, 5}

	//type inference
	//var nos = [5]int{1, 2, 3, 4, 5}
	//using :=
	nos := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nos)

	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	newNos := nos // a new copy of array is created
	newNos[0] = 999

	fmt.Println(newNos[0]) //=> 900
	fmt.Println(nos[0])    //=> 1

	//array pointer
	var nosPtr *[5]int

	//refrence
	nosPtr = &nos

	//derefrence
	fmt.Println((*nosPtr)[0])
	fmt.Println(nosPtr[0])

}
