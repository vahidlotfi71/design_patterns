package functions

import (
	"fmt"
	"slices"
	"time"
)

func FindWeekends() {
	switch time.Now().Weekday() {
	case time.Friday, time.Thursday:
		fmt.Println("Its a weekend")
	default:
		fmt.Println("it's weekday")
	}
}

func Day_nights() {
	t := time.Now()
	switch {
	case t.Hour() >= 12:
		fmt.Println("its afternoon", t.Hour())
	default:
		fmt.Println("its morning", t.Hour())
	}
}

func ConvNum() {
	fmt.Print("Enter a number between 1 and 3: ")
	var a int
	fmt.Scanln(&a)

	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("unknown number")

	}
}

func Arr() {

	var arr [3]int
	var b [4]int = [4]int{1, 2, 3}
	var c []interface{} = []interface{}{"a", "b", "c", 121, 11, 2}

	slice := make([]int, 3, 4)

	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "b", "c", "d"}

	if slices.Equal(s1, s2) {
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = (i + j)
		}
	}

	arr[1] = 12

	fmt.Println("arr len: ", len(arr))
	fmt.Println("arr: ", arr)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("twoD: ", twoD)
	fmt.Println("slice: ", slice)
}

func Map() {
	map1 := make(map[string]int)

	map1["vahid"] = 20
	map1["hasan"] = 19
	map1["mohammad"] = 18
	for key := range map1 {
		fmt.Println(key)
	}
	fmt.Println(map1)
}

func Plus(a, b int) {
	fmt.Printf("%v+%v: %v\n", a, b, a+b)
}

func VerdictFunc(numbers ...int) {
	var total int

	for _, val := range numbers {
		total += val
	}
	fmt.Println("total:", total)

}

func IntSql() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


func Recursive(n int )int{
	if n == 0 {
		return  0 
	}
	return n + Recursive(n-1)
}

func Fibo(n int)int{
	if n < 2{
		return n
	}
	return Fibo(n-1) + Fibo(n-2)
}