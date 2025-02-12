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
	f.Sum()
	f.Find_Index()
	f.Map_iter()
	f.Range1()
	num := 12
	f.Pointer1(&num)
	f.New_struct("vahid", 32)

	ract1 := f.Rect{Hight: 12, Wight: 20}
	ract1.Area()
	ract1.Prime()

	ract1.SetParameters(11, 22)
	fmt.Println(ract1)


	fmt.Println()
	


	
}

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected 
	StateError 
	StateRetrying
)

var statName  = map[ServerState]string{
	StateIdle: "state",
	StateConnected : "connected",
	StateError : "error",
	StateRetrying : "retrying",
	}

func (ss ServerState)String() string {
	return statName[ss]
}