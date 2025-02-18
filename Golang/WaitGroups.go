package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Printf("worker %d started\n", i)
	time.Sleep(time.Second * 2)
	fmt.Printf("worker %d finished\n", i)
}

func main() {
	var wg sync.WaitGroup
	for i:= range 5{
		wg.Add(1)
		go func(){
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
