package main

import "fmt"

var janelas, corredores int
var lugarAssento string

func temJanela(array []int) bool {
	for _, v := range array {
		if v == 1 {
			return true
		}
	}
	return false
}

func temCorredor(array []int) bool {
	for _, v := range array {
		if v == 1 {
			return true
		}
	}
	return false
}

func main() {
	janela := []int{}
	corredor := []int{}
	for i := 0; i < 24; i++ {
		fmt.Scan(&janelas)
		janela = append(janela, janelas)
	}
	for i := 0; i < 24; i++ {
		fmt.Scan(&corredores)
		corredor = append(corredor, corredores)
	}
	for {
		fmt.Scan(&lugarAssento)
		if lugarAssento == "corredor" {
			if !temCorredor(corredor) {
				fmt.Println("Não tem poltronas livres no corredor")
			} else {
				for i := 0; i < 24; i++ {
					if corredor[i] == 1 {
						fmt.Printf("Assento %v disponível\n", i+1)
						corredor[i] = 0
						break
					} else {
						continue
					}
				}
			}
		} else if lugarAssento == "janela" {
			if !temJanela(janela) {
				fmt.Println("Não tem poltronas livres na janela")
			} else {
				for i := 0; i < 24; i++ {
					if janela[i] == 1 {
						fmt.Printf("Assento %v disponível\n", i+1)
						janela[i] = 0
						break
					} else {
						continue
					}
				}
			}
		}
		if !temCorredor(corredor) && !temJanela(janela) {
			fmt.Println("Não tem poltronas disponíveis")
			break
		}
	}
}
