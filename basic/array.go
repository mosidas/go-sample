package basic

import (
	"fmt"
)

func Array() {
	fmt.Println("=== array ===")
	var x [3]int
	var y [3]int = [3]int{1, 2, 3}
	var z [3]int = [...]int{4, 5, 6}
	fmt.Println(x, y, z)          // [0 0 0] [1 2 3] [4 5 6]
	fmt.Println(x[0], y[1], z[2]) // 0 2 3
	a := [...]int{1, 2, 3}
	fmt.Println(a) // [1 2 3]
	b := [...]int{4, 5, 6}
	a = b             // shallow copy
	fmt.Println(a, b) // [4 5 6] [4 5 6] -> a is changed
}
