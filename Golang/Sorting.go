package main

import (
	"fmt"
	"slices"
)

func main() {
	slice1 := []string{"c", "d", "e", "a"}
	slice2 := []int{3, 1, 2, 7, 2}

	slices.Sort(slice1)
	slices.Sort(slice2)

	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)

}
