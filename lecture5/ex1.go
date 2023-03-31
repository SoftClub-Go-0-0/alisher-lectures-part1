package main

import "fmt"

func main() {
	arr := make([]int, 1000)
	var min int = 9999999999
	var max int = -9999999999
	for i := 0; i <= 1000; i++ {
		var a int
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		arr[i] = a
		
	}
	for _, v := range arr {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	if arr[0] == 0 {
		fmt.Println("Empty sequence")
	} else {
		fmt.Print(min, " ", max)
	}
}
