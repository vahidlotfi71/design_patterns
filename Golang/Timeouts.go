package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "hello world"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <- time.After(time.Second * 1):
		fmt.Println("time out1: ",msg)
	}

	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "vahid lotfi"
	}()

	select {
	case msg := <-ch2:
		fmt.Println(msg)
	case msg := <-time.After(time.Second * 1):
		fmt.Println("time out2: ",msg)
	}
}
