// Iterate over a slice
xi := []int{10, 20, 30}
for i, x := range xi {
  fmt.Println(i, x)
}
// outputs:
// 0, 10
// 1, 20
// 2, 30

xi := []int{10, 20, 30}
for _, x := range xi {
  fmt.Println(x)
}
// outputs:
// 10
// 20
// 30

xi := []int{10, 20, 30}
// for i, _ := range xi {
for i := range xi {
  fmt.Println(i)
}
// outputs:
// 0
// 1
// 2


// Iterate over a map
// Iterating over a map raises a new problem. The order is now random
hash := map[int]int{9: 10, 99: 20, 999: 30}
for k, v := range hash {
  fmt.Println(k, v)
}
// outputs, for example:
// 99 20
// 999 30
// 9 10

