
package main

import "fmt"

func main () {

	arr := make([]int, 100)

	for i := range arr{
		arr[i] = 100 - i
	}

	fmt.Print(arr)
}