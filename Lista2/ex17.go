// Exercicío 17

package main

import "fmt"

func main() {
	var tipoConsumidor string
	var m3Gastos float64
	fmt.Print("Digite o tipo do consumidor (residencial, comercial ou industrial): ")
	fmt.Scanln(&tipoConsumidor)
	fmt.Print("Digite o volume de água gasto em m3: ")
	fmt.Scanln(&m3Gastos)

	switch tipoConsumidor {
	case "residencial":
		custo := 5.00 + (0.05 * m3Gastos)
		fmt.Printf("O custo da conta de água é: R$ %.2f\n", custo)
	case "comercial":
		if m3Gastos <= 80 {
			custo := 500.00
			fmt.Printf("O custo da conta de água é: R$ %.2f\n", custo)
		} else {
			custo := 500.00 + (0.25 * (m3Gastos - 80))
			fmt.Printf("O custo da conta de água é: R$ %.2f\n", custo)
		}
	case "industrial":
		if m3Gastos <= 100 {
			custo := 800.00
			fmt.Printf("O custo da conta de água é: R$ %.2f\n", custo)
		} else {
			custo := 800.00 + (0.04 * (m3Gastos - 100))
			fmt.Printf("O custo da conta de água é: R$ %.2f\n", custo)
		}
	default:
		fmt.Println("Tipo de consumidor inválido. Por favor, escolha entre residencial, comercial ou industrial.")
	}

}
