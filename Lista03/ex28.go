package main

import (
	"fmt"
	"math"
)

var tot1, tot2 float64
var numero int

func elevadoTres(num int) int {
	return num * num * num
}

func main() {
	numero = 1
	for i := 1; i <= 51; i++ {
		if i%2 == 0 {
			tot1 += 1.0 / float64(elevadoTres(numero))
		} else {
			tot2 += 1.0 / float64(elevadoTres(numero))
		}
		numero += 2
	}
	s := tot2 - tot1
	tot3 := s * 32
	pi := math.Pow(tot3, 1.0/3.0)
	fmt.Printf("Pi é igual a %v\n", pi)
}
