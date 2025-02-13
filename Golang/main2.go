package main

import (
	"errors"
	"fmt"
)

func func1(number int) (int, error) {
	if number == 20 {
		return 0, errors.New("invalid number")
	}

	return number + 2, nil
}

var outOfRangeE = fmt.Errorf("input out of range")
var oddNumber = fmt.Errorf("input is odd number")

func GetNumber(number int) (int , error) {
	if number < 0 || number >100 {
		return 0, outOfRangeE
	}else if number % 2 == 0 {
		return 0 , oddNumber
	}
	return number,nil
}

func main() {
	fmt.Println(func1(21))
	fmt.Println(GetNumber(23))


}
