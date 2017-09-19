package main

import "fmt"

func main(){
	var array []int
	array = make([]int, 10)

	for i :=0; i < len(array); i++ {
		array[i] = i
	}

	fmt.Printf("%v", array)
}
