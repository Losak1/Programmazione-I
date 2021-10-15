/*
STAMPA UNA TABELLA PITAGORICA
*/

package main

import "fmt"

func main() {

	var d int
	fmt.Scan(&d)

	for riga := 1; riga <= d; riga++ {
		for colonna := 1; colonna <= d; colonna++ {
			fmt.Print(riga*colonna, " ")
		}
		fmt.Println()
	}
}
