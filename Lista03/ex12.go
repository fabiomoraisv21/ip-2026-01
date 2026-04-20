package main

import "fmt"

var real, sFatPares, sFatImpares, tot float64
var n int

func fat(n int) int {
	if n == 1 {
		return 1
	}
	return n * fat(n-1)
}

func main() {
	fmt.Println("Digite um número real:")
	fmt.Scan(&real)
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			sFatPares += real / float64(fat(i))
		} else {
			sFatImpares += real / float64(fat(i))
		}
	}
	tot = real - sFatImpares + sFatPares
	fmt.Println(tot)
}
