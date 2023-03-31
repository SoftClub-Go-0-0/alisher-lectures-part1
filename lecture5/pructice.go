package main

import "fmt"

type Person struct {
	Name    []string
	Surname map[string]int
	Phone   uint
	Age     uint
	Adress  string
}

func (p *Person) Sum() int {
	var cmt int
	for _,v :=range p.Surname{
		cmt+=v
	}
	return cmt
}

type Suminterface interface{
	Sum() int
}

func main() {
	var Alisher Person
	Alisher = Person{
		Name:    []string{1: "asd", 2: "sdaa"},
		Surname: map[string]int{
			"alisher" : 2,
			"behzod" : 3,
			"Salmon" : 4,
		},
		Phone:   0,
		Age:     0,
		Adress:  "",
	}
	fmt.Println(Suminterface.Sum(&Alisher))
	fmt.Println(Alisher.Sum())
}