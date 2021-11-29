/*
Scrivere un programma che legga da standard input una stringa senza spazi e, considerando l’insieme delle lettere dell'alfabeto inglese, stampi

il numero di vocali maiuscole;
il numero di vocali minuscole;
il numero di consonanti maiuscole;
il numero di consonanti minuscole.
A tal fine definire le seguenti funzioni: 'èLetteraValida(l rune) bool', 'èMaiuscola(l rune) bool', 'èVocale(l rune) bool'.

Modificare il programma (già svolto la scorsa volta) per

Utilizzare il costrutto switch per il controllo delle vocali nella funzione 'èVocale'
Utilizzare le funzioni 'unicode.IsLetter' e 'unicode.IsUpper' del package 'unicode' al posto di 'èLetteraValida' e 'èMaiuscola' rispettivamente.
*/

package main

import (
	"fmt"
	"unicode"
)

func èLetteraValida(l rune) bool {
	if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
		return true
	} else {
		return false
	}
}

func èMaiuscola(l rune) bool {
	if l >= 'A' && l <= 'Z' {
		return true
	} else {
		return false
	}
}

func èVocale(l rune) bool {
	switch unicode.ToUpper(l) {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func main() {

	var input string
	fmt.Scan(&input)

	var nVocMai, nVocMin, nConMai, nConMin int

	for _, r := range input {
		if èLetteraValida(r) {
			maiusc := èMaiuscola(r)
			vocale := èVocale(r)
			if maiusc && vocale {
				nVocMai++
			} else if maiusc {
				nConMai++
			} else if vocale {
				nVocMin++
			} else {
				nConMin++
			}
		}
	}

	fmt.Println("Numero di vocali maiuscole", nVocMai)
	fmt.Println("Numero di vocali minuscole", nVocMin)
	fmt.Println("Numero di consonanti maiuscole", nConMai)
	fmt.Println("Numero di consonanti minuscole", nConMin)
}
