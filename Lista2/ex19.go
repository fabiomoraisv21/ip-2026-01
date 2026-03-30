// Exercicío 19

package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura float64

	fmt.Println("Escolha a figura desejada:")
	fmt.Println("1 - Cone")
	fmt.Println("2 - Cilindro")
	fmt.Println("3 - Esfera")
	fmt.Scanln(&opcao)

	switch opcao {
	case 1:
		fmt.Print("Digite o raio da base do cone: ")
		fmt.Scanln(&raio)
		fmt.Print("Digite a altura do cone: ")
		fmt.Scanln(&altura)

		volumeCone := (1.0 / 3.0) * math.Pi * math.Pow(raio, 2) * altura
		areaSuperficieCone := math.Pi * raio * (raio + math.Sqrt(math.Pow(altura, 2)+math.Pow(raio, 2)))
		fmt.Printf("Volume do cone: %.2f\n", volumeCone)
		fmt.Printf("Área da superfície do cone: %.2f\n", areaSuperficieCone)

	case 2:

		fmt.Print("Digite o raio da base do cilindro: ")
		fmt.Scanln(&raio)
		fmt.Print("Digite a altura do cilindro: ")
		fmt.Scanln(&altura)

		volumeCilindro := math.Pi * math.Pow(raio, 2) * altura
		areaSuperficieCilindro := 2 * math.Pi * raio * (raio + altura)
		fmt.Printf("Volume do cilindro: %.2f\n", volumeCilindro)
		fmt.Printf("Área da superfície do cilindro: %.2f\n", areaSuperficieCilindro)

	case 3:

		fmt.Print("Digite o raio da esfera: ")
		fmt.Scanln(&raio)

		volumeEsfera := (4.0 / 3.0) * math.Pi * math.Pow(raio, 3)
		areaSuperficieEsfera := 4 * math.Pi * math.Pow(raio, 2)
		fmt.Printf("Volume da esfera: %.2f\n", volumeEsfera)
		fmt.Printf("Área da superfície da esfera: %.2f\n", areaSuperficieEsfera)

	default:
		fmt.Println("Opção inválida!")
	}
}
