package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	array := make([]string, n)
	var success float64
	var fail float64

	for i := 0; i < n; i++ {
		var a string
		fmt.Scan(&a)
		array[i] = a
	}
	for _, v := range array {
		if v == "success" {
			success++
		} else if v == "fail" {
			fail++
		}
	}
	fmt.Printf("%.2f %.2f",full(success, n), full(fail, n))
}
func full(s float64,n int) float64 {
	return float64(s * float64(100) / float64(n))
}
