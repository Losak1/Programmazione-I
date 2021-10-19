/*
Scrivere un programma che:

Genera in modo casuale un numero intero compreso nell'intervallo che va da 1 a 100,
estremi inclusi (sia numeroGenerato la variabile intera in cui viene memorizzato
il numero generato, come indicato nella consegna deve valere 1<= numeroGenerato <= 100).

Chiede iterativamente all'utente di inserire da standard input un numero intero;
ad ogni iterazione il programma controlla se il numero inserito è uguale al numero
memorizzato in numeroGenerato:

se sono uguali, il programma termina;

se sono diversi, il programma fornisce un suggerimento specificando
se il numero inserito è più alto o più basso di quello memorizzato in numeroGenerato.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(int64(time.Now().Nanosecond()))
	var numeroGenerato int = rand.Intn(100)
	var x int

	for i := 1; x != numeroGenerato; i++ {
		fmt.Print("Tentativo n°", i, ": ")
		fmt.Scan(&x)
		if x > numeroGenerato {
			fmt.Println("Troppo alto! Riprova!")
		} else if x < numeroGenerato {
			fmt.Println("Troppo basso! Riprova!")
		} else {
			fmt.Println("Hai indovinato in ", i, " tentativi!")
		}
	}
}
