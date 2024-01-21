package basic

import (
	"fmt"
)

func Printtype() {
	fmt.Println("=== type ===")
	var x int = 1
	fmt.Printf("x= %v, type= %T\n", x, x) // x= 1, type= int
	var y int8 = 2
	fmt.Printf("y= %v, type= %T\n", y, y) // y= 2, type= int8
	var z int16 = 3
	fmt.Printf("z= %v, type= %T\n", z, z) // z= 3, type= int16
	var a int32 = 4
	fmt.Printf("a= %v, type= %T\n", a, a) // a= 4, type= int32
	var b int64 = 5
	fmt.Printf("b= %v, type= %T\n", b, b) // b= 5, type= int64
	var c uint = 6
	fmt.Printf("c= %v, type= %T\n", c, c) // c= 6, type= uint
	var d uint8 = 7
	fmt.Printf("d= %v, type= %T\n", d, d) // d= 7, type= uint8
	var e uint16 = 8
	fmt.Printf("e= %v, type= %T\n", e, e) // e= 8, type= uint16
	var f uint32 = 9
	fmt.Printf("f= %v, type= %T\n", f, f) // f= 9, type= uint32
	var g uint64 = 10
	fmt.Printf("g= %v, type= %T\n", g, g) // g= 10, type= uint64
	var h uintptr = 11
	fmt.Printf("h= %v, type= %T\n", h, h) // h= 11, type= uintptr
	var i float32 = 12.1
	fmt.Printf("i= %v, type= %T\n", i, i) // i= 12.1, type= float32
	var j float64 = 13.2
	fmt.Printf("j= %v, type= %T\n", j, j) // j= 13.2, type= float64
	var k complex64 = 14 + 15i
	fmt.Printf("k= %v, type= %T\n", k, k) // k= (14+15i), type= complex64
	var l complex128 = 16 + 17i
	fmt.Printf("l= %v, type= %T\n", l, l) // l= (16+17i), type= complex128
	var m byte = 18
	fmt.Printf("m= %v, type= %T\n", m, m) // m= 18, type= uint8
	var n rune = 'あ'
	fmt.Printf("n= %v, type= %T\n", n, n) // n= 65, type= int32
	var o string = "20"
	fmt.Printf("o= %v, type= %T\n", o, o) // o= 20, type= string
	var p bool = true
	fmt.Printf("p= %v, type= %T\n", p, p) // p= true, type= bool

	// type definition
	type MyInt int
	var q MyInt = 1
	fmt.Printf("q= %v, type= %T\n", q, q) // q= 1, type= basic.MyInt

	type UtcTime string
	type JstTime string
	var t1 UtcTime = "00:00:00"
	var t2 JstTime = "09:00:00"
	//t1 = t2 // error (cannot use t2 (type JstTime) as type UtcTime in assignment)
	fmt.Printf("t1= %v, type= %T\n", t1, t1) // t1= 00:00:00, type= basic.UtcTime
	fmt.Printf("t2= %v, type= %T\n", t2, t2) // t2= 09:00:00, type= basic.JstTime
}
