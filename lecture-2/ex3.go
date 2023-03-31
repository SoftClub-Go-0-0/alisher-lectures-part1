package main

import "fmt"

func applyDiscount(x, y *int) {
	*x = *x - (*x * (*y) / 100)

}

func main() {
	for i := 1; i <= 5; i++ {
		var Price int
		var DiscountPercent int
		fmt.Scan(&Price, &DiscountPercent)
		if (Price > 0 && Price <= 100000) && (DiscountPercent >= 0 && DiscountPercent <= 100) {
			applyDiscount(&Price, &DiscountPercent)
			fmt.Println(Price)
		}
	}
}
