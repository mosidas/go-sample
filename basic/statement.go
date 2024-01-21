package basic

import (
	"fmt"
)

func Statement() {
	fmt.Println("=== statement ===")
	ifstatement()
	forstatement()

}

func ifstatement() {
	fmt.Println("--- if ---")
	// standard
	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	// true

	// declare
	if x := 1; x == 1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	// true

	// else if
	if x := 1; x == 1 {
		fmt.Println("true")
	} else if x == 2 {
		fmt.Println("false")
	} else {
		fmt.Println("else")
	}
}

func forstatement() {
	fmt.Println("--- for ---")
	// standard
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	// 0
	// 1
	// 2

	// range
	for i, x := range []int{1, 2, 3} {
		fmt.Println(i, x)
	}
	// 0 1
	// 1 2
	// 2 3

	// go is not have while. instead, use for.
	var x int = 0
	for x < 3 {
		fmt.Println(x)
		x++
	}
	// 0
	// 1
	// 2

	// break
	for i := 0; i < 3; i++ {
		if i == 1 {
			break
		}
		fmt.Println(i)
	}
	// 0

	// continue
	for i := 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}
	// 0
	// 2

	// infinite loop
	// for {
	// 	fmt.Println("forever")
	// }
	// forever
	// forever
	// :
	// :
}
