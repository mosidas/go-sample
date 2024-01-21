package basic

import (
	"fmt"
)

func Slice() {
	fmt.Println("=== slice ===")
	var x []int
	var y []int = []int{1, 2, 3, 4, 5}
	fmt.Println(x, y) // [] [1 2 3 4 5]
	// append
	x = append(x, 1)
	y = append(y, 6)
	fmt.Println(x, y) // [1] [1 2 3 4 5 6]
	// access
	fmt.Println(x[0], y[1]) // 1 2
	// length
	fmt.Println(len(x), len(y)) // 1 6
	// slice
	fmt.Println(y[1:3]) // [2 3] -> [start:end]: start <= index < end
	fmt.Println(y[1:])  // [2 3 4 5 6] -> [start:]: start <= index
	fmt.Println(y[:3])  // [1 2 3] -> [:end]: index < end
	fmt.Println(y[:])   // [1 2 3 4] -> [:]: all
	// make
	var z []int = make([]int, 3)
	fmt.Println(z) // [0 0 0]
	// make with capacity (3rd argument)
	var a []int = make([]int, 3, 5) // length: 3, capacity: 5
	fmt.Println(a)                  // [0 0 0]
	// copy
	var b []int = make([]int, len(y))
	copy(b, y) // deep copy
	fmt.Println(b)
	// mod b
	b[0] = 5
	fmt.Println(b, y) // [5 2 3 4] [1 2 3 4] -> y is not changed
}
