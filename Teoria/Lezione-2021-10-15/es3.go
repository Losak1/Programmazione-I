/*
VERIFICA SE UN NUMERO È PRIMO (NON OTTIMIZZATO)
*/

package main

import "fmt"

func main() {

	var x int
	fmt.Println("Inserisci un numero naturale:")
	fmt.Scan(&x)
	t := x
	trovato := false

	if x > 0 {
		for d := 2; d < x && !trovato; d++ {
			if x%d == 0 {
				trovato = true
			}
		}
		if trovato {
			fmt.Println(t, "è primo")
		} else {
			fmt.Println(t, "non è primo")
		}
	} else {
		fmt.Println(t, "non è naturale")
	}
}
