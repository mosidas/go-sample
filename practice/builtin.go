package practice

import (
	"fmt"
)

func init() {
	fmt.Println("imported practice package")
}

func Builtin() {
	fmt.Println("=== builtin ===")

	// println
	println("println")
	// panic
	//panic("panic")
	// recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%[1]T: %[1]s\n", r)
		}
	}()
	zero := 0
	divzero := 1 / zero
	fmt.Println(divzero)
	// init
	//init()

}
