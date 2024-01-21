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

func raisepanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	fmt.Println("do something start")
	panic("panic!")
	fmt.Println("do something end") // not called
}

func outofrange() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	// out of range
	var a []int = []int{1, 2, 3}
	fmt.Println(a[3]) // panic: runtime error: slice bounds out of range
}

func nilpointer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	// not initialize(nil pointer)
	var b *int
	fmt.Println(*b) // panic: runtime error: invalid memory address or nil pointer dereferen
}

func interfaceconversion() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	// interface conversion
	var a interface{} = 1
	var b string = a.(string) // panic: interface conversion: interface {} is int, not string
	fmt.Println(b)
}
