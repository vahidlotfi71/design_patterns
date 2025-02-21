package main

import "fmt"

type request struct {
	action   string
	value    int
	response chan int
}

func main() {

	requests := make(chan request)

	go func() {
		stat := 0
		for req := range requests {
			switch req.action {
			case "increment":
				stat += req.value
				req.response <- int(stat)
			case "get":
				req.response <- stat
			default:
				fmt.Println("Unknown action")
			}
		}
	}()

	sendRequest := func(action string, value int) int {
		respond := make(chan int)
		requests <- request{action , value, respond}
		return <- respond
	}

	fmt.Println("increment by 1: ", sendRequest("increment" , 12))
	fmt.Println("increment by 2: ", sendRequest("increment" , 12))
	fmt.Println("get current status: ", sendRequest("get",0))



}
