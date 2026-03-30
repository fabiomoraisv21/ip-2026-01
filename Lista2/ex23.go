// Exercicío 23

package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Digite a idade da pessoa: ")
	fmt.Scanln(&idade)

	switch {
	case idade < 16:
		fmt.Println("Classe eleitoral: Não-eleitor")
	case idade >= 18 && idade <= 65:
		fmt.Println("Classe eleitoral: Eleitor obrigatório")
	case (idade >= 16 && idade < 18) || idade > 65:
		fmt.Println("Classe eleitoral: Eleitor facultativo")
	default:
		fmt.Println("Idade inválida!")
	}
}
