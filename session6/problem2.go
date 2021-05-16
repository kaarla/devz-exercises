package main

import(
	"fmt"
	"sort"
)

func main() {
	fmt.Println("For $6")
	calculateGeneralCoins(6, []int{5, 2, 1,10})
	fmt.Println("For $8")
	calculateGeneralCoins(8, []int{2, 1, 10, 5})
	fmt.Println("For $10")
	calculateCoins(10)
	fmt.Println("For $5")
	calculateCoins(5)
	fmt.Println("For $2")
	calculateCoins(2)
	fmt.Println("For $1")
	calculateCoins(1)
}

func calculateCoins(n int){
	ten, five, two := 0, 0, 0
	if n >= 10 {
		ten, n = int(n / 10), n % 10
		fmt.Printf("%d coins of 10\n", ten)
	}
	if n >= 5 {
		five, n = int(n / 5), n % 5
		fmt.Printf("%d coins of 5\n", five)
	}
	if n >= 2 {
		two, n = int(n / 2), n % 2
		fmt.Printf("%d coins of 2\n", two)
	}
	if n >= 1 {
		fmt.Printf("%d coins of 1\n", n)
	}
}


func calculateGeneralCoins(n int, available []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(available)))
	coins := 0
	for _, denom := range available {
		if n >= denom {
			coins, n = int(n / denom), n % denom
			fmt.Printf("%d coins of %d\n", coins, denom)
		}
	}
}

