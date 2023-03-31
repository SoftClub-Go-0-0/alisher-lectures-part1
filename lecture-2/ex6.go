package main

import "fmt"

func isMultiple(x,y int) bool {
	var s bool
	s = y%x==0
	return s
}

func multiplesNumberInRange(a,b,c int) int {
   var cmt int
   for i:=a+1 ; i<b;i++{
	if isMultiple(c, i){
		cmt++
	} 
   }
   return cmt
}

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scanln(&b)
	fmt.Scan(&c)
	multiplesNumberInRange(a, b, c) 
	fmt.Println(multiplesNumberInRange(a, b, c) )
}