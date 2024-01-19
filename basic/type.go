package basic

import (
	"fmt"
)

func Printtype() {
	var x int = 1
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
	var k complex64 = 14 + 15i
	var l complex128 = 16 + 17i
	var m byte = 18
	var n rune = 'A'
	var o string = "20"
	var p bool = true
	fmt.Println(x, y, z, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	// 1 2 3 4 5 6 7 8 9 10 11 12.1 13.2 (14+15i) (16+17i) 18 65 string true

	// type definition
	type MyInt int
	var q MyInt = 1
	fmt.Println(q) // 1

	type UtcTime string
	type JstTime string
	var t1 UtcTime = "00:00:00"
	var t2 JstTime = "09:00:00"
	//t1 = t2 // error (cannot use t2 (type JstTime) as type UtcTime in assignment)
	fmt.Println(t1, t2)
}
