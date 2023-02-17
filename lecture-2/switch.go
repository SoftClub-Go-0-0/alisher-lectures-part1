package main

import "fmt"

func main() {
	var month uint
	fmt.Scan(&month)
	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("february")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("september")
	case 10:
		fmt.Println("october")
	case 11:
		fmt.Println("november")
	case 12:
		fmt.Println("december")
	}
}