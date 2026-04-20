package main

import "fmt"

var tot float64

func fat(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fat(n-1)
}

func main() {
	nume := 100
	denom := 0
	for i := 1; i <= 20; i++ {
		tot += float64(nume / fat(denom))
		nume -= 1
		denom += 1
	}
	fmt.Printf("O resultado dessa conta é %v\n", tot)
}
