package main

import "fmt"

var empty []int                 // an empty slice
vals := make([]int, 5) 			// create with length

withData := []int{0, 1, 2, 3, 4, 5}  // a slice pre-filled with some data
withData[1] = 5
x := withData[1] // x is now 5
newSlice := withData[2:4] // newSlice == []int{2,3}

// append Variadic Function
a := []int{1, 3}
b := append(a, 4, 2) // b == []int{1,3,4,2}

func main() {
	var numbers = make([]int,3,5)
	printSlice(numbers)
 }

 func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
 }