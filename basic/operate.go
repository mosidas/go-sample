package basic

import (
	"fmt"
)

func Operate() {
	var x int = 1
	var y int = 2
	var p bool = true

	// arithmetic operator
	fmt.Println(x + y) // 3
	fmt.Println(x - y) // -1
	fmt.Println(x * y) // 2
	fmt.Println(x / y) // 0
	fmt.Println(x % y) // 1

	// comparison operator
	fmt.Println(x == y) // false
	fmt.Println(x != y) // true
	fmt.Println(x > y)  // false
	fmt.Println(x >= y) // false
	fmt.Println(x < y)  // true
	fmt.Println(x <= y) // true

	// logical operator
	fmt.Println(p && p)   // true
	fmt.Println(p && !p)  // false
	fmt.Println(p || p)   // true
	fmt.Println(p || !p)  // true
	fmt.Println(!p || !p) // false
	fmt.Println(!p || p)  // true

	// bit operator
	fmt.Println(x & y)  // 0
	fmt.Println(x | y)  // 3
	fmt.Println(x ^ y)  // 3
	fmt.Println(x &^ y) // 1
	fmt.Println(x << y) // 4
	fmt.Println(x >> y) // 0

	// assignment operator
	x += y
	fmt.Println(x) // 3
	x -= y
	fmt.Println(x) // 1
	x *= y
	fmt.Println(x) // 2
	x /= y
	fmt.Println(x) // 1
	x %= y
	fmt.Println(x) // 1
	x &= y
	fmt.Println(x) // 0
	x |= y
	fmt.Println(x) // 2
	x ^= y
	fmt.Println(x) // 0
	x &^= y
	fmt.Println(x) // 1
	x <<= y
	fmt.Println(x) // 4
	x >>= y
	fmt.Println(x) // 1
}
