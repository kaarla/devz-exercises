package main

import(
	"fmt"
)

func main() {
	//var arr3 [4]int //no sé por qué esto no jala
	fmt.Println(getCombinations(3))
	fmt.Println(getCombinations(4))
}

func getCombinationsQuantity(n int, memo []int) int{
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 3
	}
	if memo[n] == 0 {
		memo[n] = getCombinationsQuantity(n - 1, memo) +
			getCombinationsQuantity(n -2, memo)
	}
	return memo[n]
}

func getCombinations(n int) []string{
	if n == 1 {
		return []string{"0", "1"}
	}
	if n == 2 {
		return []string{"00", "01", "10"}
	}
	
	var combinations []string
	for _, comb := range getCombinations(n - 2) {
		combinations = append(combinations, "10" + comb)
		combinations = append(combinations, "00" + comb)
		if comb[0] == '0' {
			combinations = append(combinations, "01" + comb)
		}
	}
	return combinations
}
