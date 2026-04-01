package main

import "fmt"

func main() {

	var qnumeros int

	fmt.Println("Quantos números serão analisados?")
	fmt.Scan(&qnumeros)

	numeros := make([]float64, qnumeros)

	for i := 0; i < qnumeros; i++ {
		fmt.Printf("Digite o %d° numero:", i+1)
		fmt.Scan(&numeros[i])
	}

	fmt.Println("A ordem inversa de entrada dos números é:")

	for q := qnumeros - 1; q >= 0; q-- {
		fmt.Print(numeros[q], ", ")
	}
}
