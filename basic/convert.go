package basic

import (
	"fmt"
	"strconv"
)

func Convert() {
	var x int = 1 // 32bitOS: int32, 64bitOS: int64
	var y int8 = 2
	var z int16 = 3
	var a int32 = 4
	var b int64 = 5
	var c uint = 6
	var d uint8 = 7
	var e uint16 = 8
	var f uint32 = 9
	var g uint64 = 10
	var h uintptr = 11
	var i float32 = 12.1
	var j float64 = 13.2
	var m byte = 18
	var n rune = 'A'
	var o string = "20"

	// int
	fmt.Println(int(y), int(z), int(a), int(b), int(c), int(d), int(e), int(f), int(g), int(h), int(i), int(j), int(m), int(n))

	// numeric to string
	fmt.Println(strconv.Itoa(x))

	// string to numeric
	var q, err = strconv.Atoi(o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
	// 20
}
