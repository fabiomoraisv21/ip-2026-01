package main

import "fmt"

var raio float64

func elevadoATres(num float64) float64 {
	return num * num * num
}

func main() {
	raio = 0
	for i := 0; i <= 40; i++ {
		volume := (4 * 3.14 * elevadoATres(float64(raio))) / 3
		fmt.Printf("O volume da esfera de raio %v é %v\n", raio, volume)
		raio += 0.5
	}
}
