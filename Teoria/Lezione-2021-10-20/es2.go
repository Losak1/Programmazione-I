/*
STAMPA GLI ELEMENTI DELLA SEQUENZA DI COLLATZ PARTENDO DA UN NUMERO INSERITO IN INPUT
*/

package main

import (
	"fmt"
	"unicode"
)

func maiuscuole(s string) (t string) {
	for _, r := range s {
		t += string(unicode.ToUpper(r))
	}
	return
}

func main() {
	fmt.Println(maiuscuole("Ciao MaMMaè"))
}
