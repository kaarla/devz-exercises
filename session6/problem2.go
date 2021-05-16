package main

import(
	"fmt"
)

func main() {
	fmt.Println("For $6")
	calculateCoins(6)
	fmt.Println("For $8")
	calculateCoins(8)
}

func calculateCoins(n int){
	ten, five, two := 0, 0, 0
	if n > 10 {
		ten, n = int(n / 10), n % 10
		fmt.Printf("%d coins of 10\n", ten)
	}
	if n > 5 {
		five, n = int(n / 5), n % 5
		fmt.Printf("%d coins of 5\n", five)
	}
	if n > 2 {
		two, n = int(n / 2), n % 2
		fmt.Printf("%d coins of 2\n", two)
	}
	fmt.Printf("%d coins of 1\n", n)
}
