package main

import "fmt"

func digits12(x int) (a int , b int){
	a = x /10
	b = x %10
	return a,b
}

func main() {
	for i := 1; i <= 3; i++ {
		var a int
		fmt.Scan(&a)
		digits12(a)
		fmt.Println(digits12(a))
	}
}