package main

import "fmt"

func digits1234(x int) (a,b,c,d int){
	a = x/1000
	b = x /100%10
	c = x / 10 % 10
	d = x %10
	return a,b,c,d
}
func isEven(y int)bool{
	var cmt bool
	cmt = y%2==0
	return cmt
}

func main() {
	var a int
	fmt.Scan(&a)
	_, s , _ ,e:= digits1234(a)
	isEven(s+e)
	fmt.Println(isEven(s+e))
}