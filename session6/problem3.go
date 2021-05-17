package main

import(
	"fmt"
)

func main() {
	fmt.Println("Caminos a 2,2:", getWays(2, 2))
	fmt.Println("Caminos a 2,3:", getWays(2, 3))
}

func getWays(east, north int)  int {
	if east < north {
		east, north = north, east
	}
	arr := []int{}
	for i := 0; i <= north; i++ {
		arr = append(arr, 1)
	}

	for i := 1; i < east; i++ {
		prev := 0
		for j := 0; j < north; j++ {
			arr[j] = arr[j] + prev
			prev = arr[j]
		}
	}
	return arr[north - 1]
}
