package main

import "fmt"

func main() {
	channel_1 := make(chan string, 2)

	channel_1 <- "one"
	channel_1 <- "two"
	close(channel_1)

	for elem := range channel_1 {
		fmt.Println(elem)
	}
}
