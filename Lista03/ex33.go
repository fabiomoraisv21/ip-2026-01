package main

import "fmt"

var int1, int2, resto, nSub int

func main() {
	fmt.Println("Digite o valor do primeiro número:")
	fmt.Scan(&int1)
	fmt.Println("Digite o valor do segundo número:")
	fmt.Scan(&int2)
	if int2 > int1 {
		fmt.Println("Não foi possível realizar essa divisão")
		return
	}
	for {
		resto = int1 - int2
		nSub += 1
		if resto < int2 {
			break
		}
		int1 = resto
	}

	fmt.Printf("O resto da divisão de N1 por N2 é %v e o quociente é %v\n", resto, nSub)

}
