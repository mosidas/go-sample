package basic

import (
	"fmt"
)

func Variable() {
	fmt.Println("=== variable ===")
	// single
	var x int
	x = 1
	fmt.Println(x) // 1

	// multiple
	var y, z int
	y = 2
	z = 3
	var (
		a int
		b int
	)
	a = 4
	b = 5
	fmt.Println(y, z, a, b) // 2 3 4 5
	// shorthand
	d := 4         // var a int = 4
	fmt.Println(d) // 4
	e := 4.3       // var b float64 = 4.3
	fmt.Println(e) // 4.3
	f := "hello"   // var c string = "hello"
	fmt.Println(f) // hello
}

func Constant() {
	fmt.Println("=== constant ===")
	// single
	const x int = 1
	fmt.Println(x) // 1

	// multiple
	const (
		y int = 2
		z int = 3
	)
	fmt.Println(y, z) // 2 3
	// shorthand
	const a = 4    // var a int = 4
	fmt.Println(a) // 4
}
