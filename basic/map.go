package basic

import (
	"fmt"
)

func Map() {
	fmt.Println("=== map ===")
	var x map[string]int = map[string]int{}
	var y map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	var z = map[string]int{"a": 1, "b": 2, "c": 3}
	var a map[string]int
	fmt.Println(x, y, z, a) // map[] map[a:1 b:2 c:3] map[a:1 b:2 c:3] map[]
	// add
	x["a"] = 1
	y["d"] = 4
	fmt.Println(x, y) // map[a:1] map[a:1 b:2 c:3 d:4]
	// access
	fmt.Println(x["a"], y["b"]) // 1 2
	// length
	fmt.Println(len(x), len(y)) // 1 4
	// delete
	delete(y, "d")
	fmt.Println(y) // map[a:1 b:2 c:3]
	// check
	if v, ok := y["d"]; !ok {
		fmt.Println("not found") // not found
		fmt.Println(v)           // 0
	}
}
