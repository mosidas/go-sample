package basic

import (
	"fmt"
)

func Function() {
	fmt.Println("=== function ===")
	// call
	doSomething("hello") // hello
	// return
	fmt.Println(add(1, 2)) // 3
	// multiple return
	fmt.Println(swap(1, 2)) // 2 1
	// variable arguments
	fmt.Println(sum(1, 2, 3))       // 6
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15
}

// simple function
func doSomething(s string) {
	fmt.Println(s)
}

// return
func add(x int, y int) int {
	return x + y
}

// multiple return
func swap(x int, y int) (int, int) {
	return y, x
}

// variable arguments
func sum(nums ...int) int {
	var total int = 0
	for _, num := range nums {
		total += num
	}
	return total
}
