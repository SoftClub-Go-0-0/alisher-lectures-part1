package main

import "fmt"

func main() {
	var average int
	fmt.Scan(&average)
	var absentDay int
	fmt.Scan(&absentDay)
	switch{
	case average <80 :
		fallthrough
	case absentDay > 2:
		fmt.Println("You have been excluded from the course")
	default:
		fmt.Println("You can continue the course ")
	}
}