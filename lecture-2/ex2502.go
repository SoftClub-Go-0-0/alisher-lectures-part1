package main

import "fmt"

func sumRange(x,y int) int {
var cmt int
	for i:=x;i<=y;i++{
		cmt+=i
}
return cmt
}

func main() {
	for i := 1; i <= 3; i++ {
		var a int
		var b int
		fmt.Scan(&a,&b)
		sumRange(a,b)
		fmt.Println(sumRange(a,b))
	}
}