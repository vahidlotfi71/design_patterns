package main

import (
	"fmt"
	// "time"
)

func main() {

	message := make(chan string)
	message2 := make(chan string, 2)



	go func() {
		message <- "vahid lotfi from iran"
		message2 <- "hi how are you"
		message2 <- "whats about you"
	}()

	mass := <- message
	
	for i:= 0 ; i<= len(message2) ; i++ {
		fmt.Println(<- message2)
	}
	
	fmt.Println(mass)
	// time.Sleep(time.Second)

}
