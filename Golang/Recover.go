package main

import "fmt"

func MyPanic() {
	panic("panic for test")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic status: ", r)
		}
	}()

	MyPanic()

}
