package main

import "fmt"

func sender(ch chan<- string, msg string) {
	ch <- msg
}

func recive(ch <-chan string, ch1 chan<- string) {
	msg := <-ch
	ch1 <- msg
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	sender(ch1, "Iam vahid lotfi from Iran")
	recive(ch1, ch2)

	fmt.Println(<-ch2)

}