package main

import "fmt"

var seno float64

func main() {
	for i := 0; i <= 63; i++ {
		a := float64(i) / 10
		seno = a - ((a * a * a) / 6) + ((a * a * a * a * a) / 120) - ((a * a * a * a * a * a * a) / 5040)
		fmt.Printf("%v -> %.5f\n", a, seno)
	}

}
