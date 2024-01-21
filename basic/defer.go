package basic

import (
	"fmt"
	"os"
)

func Defer() {
	fmt.Println("=== defer ===")
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	readFile()
}

func readFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// for {
	// 	// do something
	// }
	// file.Close() is called here
}
