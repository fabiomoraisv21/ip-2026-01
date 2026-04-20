package main

import "fmt"

var graos, somaGraos uint64

func main() {
	graos = 1
	for i := 1; i <= 64; i++ {
		somaGraos += graos
		graos *= 2
	}
	fmt.Println(somaGraos)

}
