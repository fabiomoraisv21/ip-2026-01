package main

import "fmt"

var numm, somaa, qntd, maior, menor, qntdImpares, somaPares, qntdPares int
var mediaa, mediaPares, pImpares float64

func main() {
	menor = 30000
	for {
		fmt.Println("Digite um número:")
		fmt.Scan(&numm)
		if numm == 30000 {
			break
		}
		qntd += 1
		somaa += numm
		mediaa = float64(somaa) / float64(qntd)
		if numm > maior {
			maior = numm
		}
		if numm < menor {
			menor = numm
		}
		if numm%2 == 0 {
			qntdPares += 1
			somaPares += numm
		}
		mediaPares = float64(somaPares) / float64(qntdPares)
		if numm%2 != 0 {
			qntdImpares += 1
		}
	}
	pImpares := (float64(qntdImpares) / float64(qntd)) * 100
	fmt.Printf("A soma dos números digitados é %v\n", somaa)
	fmt.Printf("A quantidade de números digitados é %v\n", qntd)
	fmt.Printf("A média dos números digitados é %.2f\n", mediaa)
	fmt.Printf("O maior número digitado foi %v\n", maior)
	fmt.Printf("O menor número digitado foi %v\n", menor)
	fmt.Printf("A média dos números pares é %.2f\n", mediaPares)
	fmt.Printf("A porcentagem dos números ímpares entre todos os números digitados é %.1f%%\n", pImpares)
}
