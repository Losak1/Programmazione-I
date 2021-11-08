/*
ESEMPI CON STRUTTURE
*/

package main

import "fmt"

func main() {

	var studente struct {
		nome, cognome string
		annonascita   int
	}

	studente.nome = "Harvey"
	studente.cognome = "Specter"
	studente.annonascita = 1981

	fmt.Println(studente)
}
