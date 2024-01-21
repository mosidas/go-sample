package basic

import (
	"fmt"
)

func Format() {
	fmt.Println("=== format ===")
	fmt.Printf("Hello World.")
	fmt.Printf("Hello %v\n", "World") // Hello World
	var x int = 1
	var y rune = 'あ'
	var p = struct {
		X int
		Y int
	}{X: 1, Y: 2}

	// format specifier

	// %v is default
	fmt.Printf("%v\n", x) // 1
	fmt.Printf("%v\n", y) // 12354
	fmt.Printf("%v\n", p) // {1 2}
	// %#v is Go syntax
	fmt.Printf("%#v\n", p) // struct { X int; Y int }{X:1, Y:2}
	fmt.Printf("%#v\n", y) // 12354

	fmt.Printf("%d\n", 123)  // 123
	fmt.Printf("%f\n", 3.14) // 3.140000
	fmt.Printf("%e\n", 3.14) // 3.140000e+00
	fmt.Printf("%g\n", 3.14) // 3.14 %g is %e for large exponents, %f otherwise

	// %t is boolean
	fmt.Printf("%t\n", true) // true
	// %s is string or slice
	fmt.Printf("%s\n", "Hello") // Hello
	// %q is double-quoted string
	fmt.Printf("%q\n", "Hello") // "Hello"

	// %b is base 2
	fmt.Printf("%b\n", 65) // 1000001
	// %o is base 8
	fmt.Printf("%o\n", 65) // 101
	// %x is base 16
	fmt.Printf("%x\n", 65) // 41
	// %U is Unicode format
	fmt.Printf("%U\n", 65) // U+0041
	fmt.Printf("%U\n", y)  // U+3042
	// %c is Unicode character (rune)
	fmt.Printf("%c\n", 65) // A
	fmt.Printf("%c\n", y)  // あ

	// %p is pointer or channel
	fmt.Printf("%p\n", &x) // 0xc000118000
	// %T is type
	fmt.Printf("%T\n", p) // struct { X int; Y int}

	z := "Hello"
	pi := 3.14159265
	// println is print with newline using %v
	fmt.Println(z, pi) // Hello 3.14159265
}
