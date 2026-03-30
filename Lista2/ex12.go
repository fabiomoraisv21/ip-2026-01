// Exercicío 12

package main

import "fmt"

func main() {

	var idade int
	fmt.Println("Digite a idade de uma pessoa:")
	fmt.Scan(&idade)

	if idade >= 0 && idade <= 2 {
		fmt.Printf("A pessoa é um recém-nascido")
	} else if idade >= 3 && idade <= 11 {
		fmt.Printf("A pessoa é uma criança")
	} else if idade >= 12 && idade <= 19 {
		fmt.Printf("A pessoa é um adolescente")
	} else if idade >= 20 && idade <= 55 {
		fmt.Printf("A pessoa é um adulto")
	} else if idade > 55 {
		fmt.Printf("A pessoa é um idoso")
	}
}
