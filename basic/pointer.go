package basic

import (
	"fmt"
)

func Pointer() {
	fmt.Println("=== pointer ===")
	// declare
	var x int = 1
	var y *int = &x
	fmt.Println(x, y) // 1 0xcXXXXXXXXX(=address of x)
	// initialize
	var z *int = new(int)
	fmt.Println(z) // 0xcXXXXXXXXX(=address of z)
	// dereference
	fmt.Println(*y, *z) // 1 0
	// nil
	var p *int = nil
	fmt.Println(p) // <nil>
	// change value
	*y = 2
	fmt.Println(x, *y) // 2 2
}
