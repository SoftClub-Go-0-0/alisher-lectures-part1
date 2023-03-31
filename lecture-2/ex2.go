package main

import "fmt"

func eqTwo(a,b,c int) bool {
var cmt bool
 cmt =  a ==b || a ==c || b ==c
 return cmt
}

func main() {
	for i := 1; i <= 3; i++ {
		var a int
		var b int
		var c int
		fmt.Scan(&a,&b,&c)
		eqTwo(a,b,c)
		fmt.Println(eqTwo(a,b,c))

	}
}