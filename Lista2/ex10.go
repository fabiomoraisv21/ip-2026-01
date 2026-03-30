// Exercicío 10

package main

import "fmt"

func main() {
	var destino int
	var retorno int
	fmt.Print("Digite o destino (1-4): ")
	fmt.Scan(&destino)
	fmt.Print("Digite 1 para ida e volta, 2 para só ida: ")
	fmt.Scan(&retorno)

	if retorno != 1 && retorno != 2 {
		fmt.Println("Valor inválido para retorno!")
	} else if destino == 1 && retorno == 1 {
		fmt.Println("O preço da passagem é R$ 900,00")
	} else if destino == 1 && retorno == 2 {
		fmt.Println("O preço da passagem é R$ 500,00")
	} else if destino == 2 && retorno == 1 {
		fmt.Println("O preço da passagem é R$ 650,00")
	} else if destino == 2 && retorno == 2 {
		fmt.Println("O preço da passagem é R$ 350,00")
	} else if destino == 3 && retorno == 1 {
		fmt.Println("O preço da passagem é R$ 600,00")
	} else if destino == 3 && retorno == 2 {
		fmt.Println("O preço da passagem é R$ 350,00")
	} else if destino == 4 && retorno == 1 {
		fmt.Println("O preço da passagem é R$ 550,00")
	} else if destino == 4 && retorno == 2 {
		fmt.Println("O preço da passagem é R$ 300,00")
	} else {
		fmt.Println("Valor inválido para destino!")
	}
}
