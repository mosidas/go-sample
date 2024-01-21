package basic

import (
	"fmt"
)

func PanicAndRecover() {
	fmt.Println("=== panic ===")
	// start
	fmt.Println("start")

	// panic
	raisepanic()

	// out of range
	outofrange()

	// nil pointer
	nilpointer()

	// interface conversion
	interfaceconversion()

	// end
	fmt.Println("end") // not called
}

func printRecover(err interface{}) {
	if err != nil {
		fmt.Println("recover:", err)
	}
}

func raisepanic() {
	defer func() {
		printRecover(recover())
	}()
	fmt.Println("raisepanic start")
	panic("panic!")
	fmt.Println("raisepanic end") // not called
}

func outofrange() {
	defer func() {
		printRecover(recover())
	}()
	// out of range
	var a []int = []int{1, 2, 3}
	fmt.Println(a[3]) // panic: runtime error: slice bounds out of range
}

func nilpointer() {
	defer func() {
		printRecover(recover())
	}()
	// not initialize(nil pointer)
	var b *int
	fmt.Println(*b) // panic: runtime error: invalid memory address or nil pointer dereferen
}

func interfaceconversion() {
	defer func() {
		printRecover(recover())
	}()
	// interface conversion
	var a interface{} = 1
	var b string = a.(string) // panic: interface conversion: interface {} is int, not string
	fmt.Println(b)
}
