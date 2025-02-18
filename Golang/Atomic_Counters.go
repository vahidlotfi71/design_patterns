package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
func main() {
	var numb atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 5000; i++ {

		wg.Add(1)
		
		go func() {
			for j := 0; j < 10000000; j++ {
				numb.Add(1)
			}

		defer wg.Done()

		}()
	
	}

	wg.Wait()
	fmt.Println("numb: ", numb.Load())

}
