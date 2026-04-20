package main

import "fmt"

var natural1, natural2, primos, nPrimos int

func divisores(n int) int {
	qntdDivisores := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			qntdDivisores += 1
		}
	}
	if qntdDivisores == 2 {
		return 1
	}
	return 0
}

func main() {
	fmt.Println("Digite dois números naturais:")
	fmt.Scan(&natural1, &natural2)
	for i := natural1; i <= natural2; i++ {
		nPrimos += divisores(i)
	}
	fmt.Println(nPrimos)
}
