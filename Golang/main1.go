package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (a *base) describe() string {
	return fmt.Sprintf("base with num %d", a.num)
}

func (c *container) Printcontainer(){
	fmt.Printf("base with num %d and string is %s\n", c.num,c.str)

}






func main() {
	num1 := base{num: 12}
	container1 := container{base: base{num: 10},
							str: "vahid lotfi",}

	container1.Printcontainer()
	fmt.Println(container1)
	fmt.Println(num1.describe())

	type describer interface {
		describe()  string
	}
	
	
	var d describer = container1

	fmt.Println("describe: " , d.describe())
}
