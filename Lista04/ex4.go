package main

import "fmt"

func main() {
	// 1. Usamos o make para criar um slice de 10 posições
	A := make([]int, 10)

	// 2. Criamos um mapa para contar as repetições.
	// A chave (int) será o número do array, e o valor (int) será a contagem.
	contagem := make(map[int]int)

	fmt.Println("Digite 10 números inteiros:")

	// 3. Lendo os 10 números
	for i := range A {
		fmt.Printf("Posição %d: ", i+1)
		fmt.Scan(&A[i])
		
		// O Pulo do Gato: 
		// Toda vez que lemos um número, vamos no mapa e somamos +1 na contagem dele.
		// Se o número ainda não existe no mapa, o Go cria ele com valor 0 e depois soma 1.
		contagem[A[i]]++ 
	}

	fmt.Println("\n--- Resultado ---")
	
	// Variável flag (bandeira) para caso nenhum número se repita
	temRepetido := false

	// 4. Percorrendo o mapa para mostrar apenas os repetidos
	// No range de um mapa, ele devolve a 'chave' (o número) e o 'valor' (a contagem)
	for numero, vezes := range contagem {
		if vezes > 1 {
			fmt.Printf("O número %d repete-se %d vezes.\n", numero, vezes)
			temRepetido = true
		}
	}

	// 5. Se a flag continuar falsa, avisamos o utilizador
	if !temRepetido {
		fmt.Println("Não há nenhum número repetido no vetor.")
	}
}