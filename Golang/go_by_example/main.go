package main

import (
	"fmt"
	f "go_by_example/functions"
)

func main() {
	f.FindWeekends()
	f.Day_nights()
	// f.ConvNum()
	f.Arr()
	f.Map()
	f.Plus(12, 11)
	f.VerdictFunc(12, 11, 11, 76564, 121221, 13325)
	arr := []int{1, 23, 4, 5, 6, 7}
	f.VerdictFunc(arr...)

	nextInt := f.IntSql()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	recursive := f.Recursive(10)
	fmt.Println(recursive)

	fibo := f.Fibo(12)
	fmt.Println(fibo)

}
