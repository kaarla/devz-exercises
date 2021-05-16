package main

import("fmt")


func findPascalLine(n int) []uint64{
    result := []uint64{1}
    switch(n){
    	case 0: return []uint64{1}
    	case 1: return []uint64{1,1}
    	default:
    	    base := findPascalLine(n-1)
    	    for i := 0; i < len(base) -1; i++{
    	    	result = append(result, base[i] + base[i+1])
    	    }
    	    return append(result, 1)
    } 
    
}

func main(){
    //for i := 0; i < 21; i++{
    	fmt.Println(findPascalLine(500))
    //}	
}
