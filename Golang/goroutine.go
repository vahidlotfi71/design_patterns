package main

import (
	"fmt"
	"time"
)

func repete(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, ":", i)
	}

}

func main() {
	for i := range 6 {
		fmt.Println("number: ", i)
	}

	go repete("vahid")

	time.Sleep(time.Second)

}
