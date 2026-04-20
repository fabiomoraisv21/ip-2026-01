package main

import "fmt"

var idade, entrada, mais50, qntdAlturas, pessoas, pessoasMenos40Kg int
var altura, peso, mediaAlturas, porcentagem, somaAlturas float64

func main() {
	for {
		pessoas += 1
		fmt.Println("Qual é sua idade?")
		fmt.Scan(&idade)
		if idade > 50 {
			mais50 += 1
		}
		fmt.Println("Qual é sua altura?")
		fmt.Scan(&altura)
		fmt.Println("Qual é seu peso?")
		fmt.Scan(&peso)
		if peso < 40 {
			pessoasMenos40Kg += 1
		}
		if idade >= 10 && idade <= 20 {
			somaAlturas += altura
			qntdAlturas += 1
		}
		mediaAlturas = float64(somaAlturas) / float64(qntdAlturas)
		porcentagem = (float64(pessoasMenos40Kg) / float64(pessoas)) * 100
		fmt.Println(("Quer continuar digitando dados?\n1- SIM\nOutro valor- NÃO"))
		fmt.Scan(&entrada)
		if entrada != 1 {
			break
		}
	}
	fmt.Printf("A quantidade de pessoas com idade superior a 50 anos é %v\n", mais50)
	fmt.Printf("A média das alturas das pessoas com idade entre 10 e 20 anos é %.2f\n", mediaAlturas)
	fmt.Printf("A porcentagem de pessoas com peso inferior a 40 quilos entre todas as pessoas analisadas é %.1f%%\n", porcentagem)
}
