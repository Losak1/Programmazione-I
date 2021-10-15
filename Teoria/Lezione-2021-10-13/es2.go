/*
VERIFICA SE UN NUMERO È PARI O DISPARI
*/

package main

import "fmt"

func main() {

	var x int
	fmt.Println("Inserisci un numero non negativo")
	fmt.Scan(&x)

	if x >= 0 {
		if x%2 == 0 {
			fmt.Println("Il numero è pari")
		} else {
			fmt.Println("Il numero è dispari")
		}
	} else {
		fmt.Println("Il numero è negativo")
	}
}
