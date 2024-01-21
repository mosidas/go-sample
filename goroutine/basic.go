package goroutine

import (
	"fmt"
	"time"
)

func Goroutine() {
	fmt.Println("=== goroutine ===")
	// parallel execution printA() and printM()
	go printA()
	printM()
	fmt.Println()
}

func printA() {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
		time.Sleep(10 * time.Millisecond)
	}
}

func printM() {
	for i := 0; i < 10; i++ {
		fmt.Print("M")
		time.Sleep(10 * time.Millisecond)
	}
}
