package library

import (
	"fmt"
	"strconv"
)

func Strconv() {
	fmt.Println("=== strconv ===")
	f := "%T %v\n"
	// string <-> int
	s := strconv.Itoa(123)
	fmt.Printf(f, s, s) // string 123

	i, _ := strconv.Atoi("123")
	fmt.Printf(f, i, i) // int 123
	j, _ := strconv.ParseInt("123", 10, 64)
	fmt.Printf(f, j, j) // int64 123

	// string <-> float
	str := strconv.FormatFloat(1.23456, 'f', 4, 64)
	fmt.Printf(f, str, str) // string 1.2346

	flt, _ := strconv.ParseFloat("1.23", 64)
	fmt.Printf(f, flt, flt) // float64 1.23
	flt2, _ := strconv.ParseFloat("1.23", 64)
	fmt.Printf(f, flt2, flt2) // float64 1.23
}
