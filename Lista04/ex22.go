package main

import "fmt"

var codigos, opcao, contaEntrada, valorsaque, valorDeposito int
var saldos, somaSaldos float64

func achouConta(array []int, value int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func qualConta(array []int, value int) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return 0
}

func main() {
	vetorCodigos := []int{}
	vetorSaldos := []float64{}
	for i := 0; i < 10; i++ {
		fmt.Scan(&codigos)
		vetorCodigos = append(vetorCodigos, codigos)
	}
	for i := 0; i < 10; i++ {
		fmt.Scan(&saldos)
		vetorSaldos = append(vetorSaldos, saldos)
	}
	for {
		fmt.Println("1. Efetuar depósito")
		fmt.Println("2. Efetuar saque")
		fmt.Println("3. Consultar o ativo bancário (ou seja, o somatório dos saldos de todos os clientes)")
		fmt.Println("4. Finalizar o programa")
		fmt.Scan(&opcao)
		if opcao == 1 {
			fmt.Println("Código da conta:")
			fmt.Scan(&contaEntrada)
			fmt.Println("Valor do depósito:")
			fmt.Scan(&valorDeposito)
			if !achouConta(vetorCodigos, contaEntrada) {
				fmt.Println("Conta não encontrada")
				continue
			} else {
				indiceConta := qualConta(vetorCodigos, contaEntrada)
				vetorSaldos[indiceConta] += float64(valorDeposito)
				fmt.Println(vetorSaldos[indiceConta])
				continue
			}
		} else if opcao == 2 {
			fmt.Println("Código da conta:")
			fmt.Scan(&contaEntrada)
			fmt.Println("Valor do saque:")
			fmt.Scan(&valorsaque)
			if !achouConta(vetorCodigos, contaEntrada) {
				fmt.Println("Conta não encontrada")
				continue
			} else {
				indiceConta := qualConta(vetorCodigos, contaEntrada)
				if vetorSaldos[indiceConta] >= float64(valorsaque) {
					vetorSaldos[indiceConta] -= float64(valorsaque)
					fmt.Println(vetorSaldos[indiceConta])
					continue
				} else {
					fmt.Println("Saldo insuficiente")
					continue
				}
			}
		} else if opcao == 3 {
			somaSaldos = 0
			for _, v := range vetorSaldos {
				somaSaldos += v
			}
		} else {
			break
		}
	}
}
