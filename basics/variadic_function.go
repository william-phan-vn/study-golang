package main

import "fmt"

// The variadic parameter must be the last parameter of the function.
func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)

    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            return
        }
    }

    fmt.Println(num, "not found in ", nums)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
    fmt.Printf("type of nums is %T\n", value)
	return append(value, slice...)
}

func main() {
    list := []int{1, 2, 3}
    // find(1, list...) // "find" defined as shown above
    fmt.Println(PrependItems(list, 1))
}