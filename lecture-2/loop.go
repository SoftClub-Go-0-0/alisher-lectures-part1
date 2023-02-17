package main

import "fmt"

func main() {
	fmt.Print("Enter the number of tasks you've done : ")
	var numberoftask int
	fmt.Scan(&numberoftask)
	fmt.Println("enter scores you got for each",numberoftask,"tasks")
	var score , sum float64
	for i:= 1; i<=numberoftask;i++{
		fmt.Print(i,"task: ")
		fmt.Scan(&score)
		sum+=score
	}
	fmt.Println("The average score for all your tasks is ",sum / float64(numberoftask))
}