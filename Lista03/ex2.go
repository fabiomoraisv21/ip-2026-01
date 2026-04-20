package main

import "fmt"

var media float64
var soma, qntd_pares int

func main() {
	for i := 50; i <= 70; i++ {
		if i%2 != 0 {
			continue
		}
		soma += i
		qntd_pares += 1
	}
	media = float64(soma) / float64(qntd_pares)
	fmt.Printf("O resultado da soma dos valores entre 50 e 70 é %v e a média é %v\n", soma, media)
}
