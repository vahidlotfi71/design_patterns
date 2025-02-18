package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println("received",msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi vahid 20"

	select{
	case ch2 <- msg:
		fmt.Println("message sent: ",msg)
	default:
		fmt.Println("no message sen")
	}



}