/*
FUNZIONE CHE INCREMENTA UNA VARIABILE
*/

package main

import "fmt"

func inc(x *int) {
	*x++
}

func main() {

	var x int = 5
	inc(&x)
	fmt.Println(x)
}
