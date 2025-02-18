package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start")

	ticker := time.NewTicker(time.Microsecond * 200)

	go func() {
		for {
			timeT := <-ticker.C
			fmt.Println("ticker at: ", timeT)

		}
	}()

	time.Sleep(time.Second * 2)
	ticker.Stop()
	fmt.Println("ticker closed")

}
