package main

import "os"

func main() {
	panic("a problem acured")

	_, err := os.Create("/asasatest.go")
	if err != nil {
		panic(err)
	}
}
