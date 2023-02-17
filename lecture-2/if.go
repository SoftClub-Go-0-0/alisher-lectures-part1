package main

import "fmt"

func main() {
	var averagescore int
	fmt.Scan(&averagescore)
	if averagescore < 80 {
		fmt.Println("soory, you can not continue the course ")

	} else if averagescore > 100 {
		fmt.Println("Hmm...l think you've made mistake. enter the number from 0 to 100")

	}else {
		fmt.Println("Congrats! you passed the next stage.")
	}
}