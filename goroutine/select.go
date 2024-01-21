package goroutine

import (
	"fmt"
	"time"
)

func Select() {
	fmt.Println("=== select ===")
	// make channel
	ch1 := make(chan int)
	ch2 := make(chan int)

	// send data to channel
	go sendData2(ch1)
	go sendData2(ch2)

	// receive data from channel
	getData2(ch1, ch2)
}

func sendData2(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(10 * time.Millisecond)
	}

	close(ch)
}

func getData2(ch1 <-chan int, ch2 <-chan int) {
	c1 := false
	c2 := false
	for {
		select {
		case i, ok := <-ch1:
			if ok {
				fmt.Println("ch1 received:", i)
			} else {
				fmt.Println("ch1 closed")
				c1 = true
			}
		case i, ok := <-ch2:
			if ok {
				fmt.Println("ch2 received:", i)
			} else {
				fmt.Println("ch2 closed")
				c2 = true
			}
		}
		if c1 && c2 {
			return
		}
	}
}
