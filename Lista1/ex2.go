// 2. Arrecadação de Jogos
package main

import "fmt"

func main() {
	var total int
	var pop, ger, arq, cad float64

	fmt.Println("Qual o número total de pessoas que irão assistir ao jogo?")
	fmt.Scan(&total)

	fmt.Println("Qual a porcentagem das categorias? Escreva: popular, geral, arquibancada e cadeiras")
	fmt.Scan(&pop, &ger, &arq, &cad)

	arrec := float64(total) / 100 * (pop + ger*5 + arq*10 + cad*20)

	fmt.Println("A arrecadação total do jogo será de R$", arrec)
}
