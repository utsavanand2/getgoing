package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Printf("result of adding %d and %d is %d\n", a, b, add(a, b))
}
