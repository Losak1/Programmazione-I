/*
ESEMPI CON STRUTTURE
*/

package main

import "fmt"

type studente struct {
	nome, cognome string
	annonascita   int
}

func main() {

	var s1, s2 studente

	s1.nome = "Harvey"
	s1.cognome = "Specter"
	s1.annonascita = 1972

	s2.nome = "Louis"
	s2.cognome = "Litt"
	s2.annonascita = 1970

	fmt.Println(s1, s2)
}
