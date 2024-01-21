package basic

import (
	"fmt"
)

func Operate() {
	fmt.Println("=== operate ===")
	var x int = 3
	var y int = 6

	// arithmetic operator
	fmt.Printf("%v + %v = %v\n", x, y, x+y)  // 3 + 6 = 9
	fmt.Printf("%v - %v = %v\n", x, y, x-y)  // 3 - 6 = -3
	fmt.Printf("%v * %v = %v\n", x, y, x*y)  // 3 * 6 = 18
	fmt.Printf("%v / %v = %v\n", x, y, x/y)  // 3 / 6 = 0
	fmt.Printf("%v %% %v = %v\n", x, y, x%y) // 3 % 6 = 3

	// comparison operator
	fmt.Printf("%v == %v = %v\n", x, y, x == y) // 3 == 6 = false
	fmt.Printf("%v != %v = %v\n", x, y, x != y) // 3 != 6 = true
	fmt.Printf("%v > %v = %v\n", x, y, x > y)   // 3 > 6 = false
	fmt.Printf("%v >= %v = %v\n", x, y, x >= y) // 3 >= 6 = false
	fmt.Printf("%v < %v = %v\n", x, y, x < y)   // 3 < 6 = true
	fmt.Printf("%v <= %v = %v\n", x, y, x <= y) // 3 <= 6 = true

	var p bool = true
	var q bool = false
	// logical operator
	fmt.Printf("%v && %v = %v\n", p, q, p && q) // true && false = false
	fmt.Printf("%v || %v = %v\n", p, q, p || q) // true || false = true
	fmt.Printf("!%v = %v\n", p, !p)             // !true = false

	// bit operator
	fmt.Printf("%v & %v = %v\n", x, y, x&y)   // 3 & 6 = 2 : 0011 & 0110 = 0010
	fmt.Printf("%v | %v = %v\n", x, y, x|y)   // 3 | 6 = 7 : 0011 | 0110 = 0111
	fmt.Printf("%v ^ %v = %v\n", x, y, x^y)   // 3 ^ 6 = 5 : 0011 ^ 0110 = 0101
	fmt.Printf("%v &^ %v = %v\n", x, y, x&^y) // 3 &^ 6 = 1 : 0011 &^ 0110 = 0001
	fmt.Printf("%v << %v = %v\n", x, y, x<<y) // 3 << 6 = 192 : 0011 << 6 = 11000000
	fmt.Printf("%v >> %v = %v\n", x, y, x>>y) // 3 >> 6 = 0 : 0011 >> 6 = 000000

	// assignment operator
	x += y
	fmt.Printf("x += y -> %v\n", x) // 9
	x -= y
	fmt.Printf("x -= y -> %v\n", x) // 3
	x *= y
	fmt.Printf("x *= y -> %v\n", x) // 18
	x /= y
	fmt.Printf("x /= y -> %v\n", x) // 3
	x %= y
	fmt.Printf("x %%= y -> %v\n", x) // 3
	x &= y
	fmt.Printf("x &= y -> %v\n", x) // 2 : 0011 & 0110 = 0010
	x |= y
	fmt.Printf("x |= y -> %v\n", x) // 6 : 0010 | 0110 = 0110
	x ^= y
	fmt.Printf("x ^= y -> %v\n", x) // 4 : 0110 ^ 0110 = 0000
	x = 3
	x <<= y
	fmt.Printf("x <<= y -> %v\n", x) // 192 : 0011 << 6 = 11000000
	x >>= y
	fmt.Printf("x >>= y -> %v\n", x) // 3 : 11000000 >> 6 = 00000011
	x &^= y
	fmt.Printf("x &^= y -> %v\n", x) // 1 : 00000011 &^ 0110 = 00000001
}
