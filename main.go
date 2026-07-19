package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	go func() {
		time.Sleep(30 * time.Millisecond)
		ch <- "slow"
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch <- "fast"
	}()
	// select to receive first, with 100ms timeout
	select {
	case v := <-ch:
		fmt.Print(v)
	case <-time.After(time.Duration(time.Millisecond * 100)):
		fmt.Print("timeout")

	}
	_ = fmt.Print
}
