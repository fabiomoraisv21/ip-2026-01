package main

import "fmt"

var precoIngresso, ingressos, despeza, maiorLucro, lucro, precoMaiorLucro, ingressosMaiorLucro float64

func main() {
	fmt.Println("LUCRO / PREÇO DO INGRESSO")
	precoIngresso = 6
	ingressos = 130
	despeza = 300
	for {
		if precoIngresso < 1 {
			break
		}
		lucro = (precoIngresso * ingressos) - despeza
		if lucro > maiorLucro {
			maiorLucro = lucro
			precoMaiorLucro = precoIngresso
			ingressosMaiorLucro = ingressos
		}
		fmt.Printf("%.1f, %.1f\n", lucro, precoIngresso)
		precoIngresso -= 0.6
		ingressos += 30
	}
	fmt.Printf("O maior lucro esperado é %.1f, com a venda de %v ingressos a %.1f reais\n", maiorLucro, ingressosMaiorLucro, precoMaiorLucro)
}
