package main

import "fmt"

var numeroBoi, pesoBoi, maiorPeso, menorPeso, nLeve, nPesado int

func main() {
	menorPeso = 100000000
	for i := 0; i < 3; i++ {
		fmt.Println("Digite o número de identificação do boi:")
		fmt.Scan(&numeroBoi)
		fmt.Println("Digite o peso do boi em quilos:")
		fmt.Scan(&pesoBoi)
		if pesoBoi > maiorPeso {
			maiorPeso = pesoBoi
			nPesado = numeroBoi
		}
		if pesoBoi < menorPeso {
			menorPeso = pesoBoi
			nLeve = numeroBoi
		}
	}
	fmt.Printf("O boi mais pesado é o de número %v com %v quilos\n", nPesado, maiorPeso)
	fmt.Printf("O boi mais leve é o de número %v com %v quilos\n", nLeve, menorPeso)
}
