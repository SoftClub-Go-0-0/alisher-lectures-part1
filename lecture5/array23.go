package main 

import "fmt"

func main(){
	var n int 
	fmt.Scan(&n)
	var k int
	var l int
	fmt.Scan(&k,&l)
	array23 := make([]float64,n)
	for i:=0 ; i<n ; i++{
		fmt.Scan(&array23[i])
	}
	var a float64
	for i:=0 ; i<k ; i++{
		a += array23[i]
	}
	for i:=l ; i<n ; i++{
		a += array23[i]
	}
	fmt.Print(a/float64(len(array23)))

}