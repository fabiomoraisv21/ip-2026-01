package main

import (
	"fmt"
	"math"
)

func fatorial(n int) float64 {
	if n == 1 || n == 0 {
		return 1
	}
	return float64(n) * fatorial(n-1)
}

var expoentee, denominadorr int
var somaSub, somaSom, entrada, cosseno float64

func main() {
	cosseno = 1
	fmt.Println("Digite o valor de X:")
	fmt.Scan(&entrada)
	entrada = entrada * math.Pi / 180
	denominadorr = 2
	expoentee = 2
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			cosseno += (math.Pow(entrada, float64(expoentee))) / float64(fatorial(denominadorr))
		} else {
			cosseno -= (math.Pow(entrada, float64(expoentee))) / float64(fatorial(denominadorr))
		}
		expoentee += 2
		denominadorr += 2
	}
	fmt.Println("O valor do cosseno é", cosseno)
	diferenca := cosseno - math.Cos(entrada)
	fmt.Println("A diferença é", diferenca)
}
