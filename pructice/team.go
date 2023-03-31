package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i:=1 ; i<=n ; i++{
		slice := make([]int,3)
		for k,_ := range slice {
			var num int
			fmt.Scan(&num)
			slice[k]=num
		}
	}
}