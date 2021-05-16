package main

import (
    "fmt"  
)

func main(){
    arr := []int{2, 7, 1, 5}
    target := 9
    fmt.Println(findIndexes(arr, target))    
	
    arr2 := []int{3, -1, 0, 1}
    target2 := 0
    fmt.Println(findIndexes(arr2, target2))    
	    
}

func findIndexes(numbers []int, target int) []int{
    mapp := map[int]int{}
    for i, valuei := range(numbers){	
    	subs := target - valuei
    	index, ok := mapp[subs]
    	if ok {
    	    return []int{index, i}
    	}
    	mapp[valuei] = i
    }
    return []int{}
}
