package main

import (
	"design_patterns/creational"
	"fmt"
)

func main() {
	singlton1 := creational.GetInstance()
	singlton2 := creational.GetInstance()
	fmt.Println(singlton1 == singlton2)
}