package goroutine

import (
	"fmt"
)

func Channel() {
	fmt.Println("=== channel ===")
	// make channel
	ch := make(chan int)

	// send data to channel
	go sendData(ch)

	// receive data from channel
	getData(ch)
}

func sendData(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)
}

func getData(ch <-chan int) {
	for {
		i, ok := <-ch
		if ok {
			fmt.Println("received:", i)
		} else {
			break
		}
	}
}

func Deadlock() {
	chA := make(chan string)
	chB := make(chan string)
	defer close(chA)
	defer close(chB)

	go func() {
		v := "a"
		chA <- v    // send (blocking)
		v2 := <-chB // receive
		fmt.Println(v, v2)
	}()

	v := "b"
	chB <- v    // send (blocking)
	v2 := <-chA // receive
	fmt.Println(v, v2)
}
