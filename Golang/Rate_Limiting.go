package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	tiker := time.Tick(time.Second)

	for i := range 5 {
		time := <-tiker
		fmt.Printf("%d loop time:%v\n", i, time)
	}
}
