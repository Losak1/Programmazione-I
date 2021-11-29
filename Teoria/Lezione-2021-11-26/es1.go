/*
TESTO DEL PROBLEMA
*/

package main

import "fmt"

func main() {

	f := func(x int) int {
		return x * x
	}

	var x int

	fmt.Scan(&x)

	fmt.Println(f(x))
}
