package main

import(
	"fmt"
)

func main() {
	//var arr3 [4]int //no sé por qué esto no jala
	arr3 := []int{0,0,0,0}
	fmt.Println("comb de 3:", getCombinations(3, arr3))
	arr4 := []int{0,0,0,0,0}
	fmt.Println("comb de 4:", getCombinations(4, arr4))
	arr5 := []int{0,0,0,0,0,0}
	fmt.Println("comb de 4:", getCombinations(5, arr5))

}

func getCombinations(n int, memo []int) int{
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 3
	}
	if memo[n] == 0 {
		memo[n] = getCombinations(n - 1, memo) +
			getCombinations(n -2, memo)
	}
	return memo[n]
}
