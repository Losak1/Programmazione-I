/*
Scrivere un unico programma che:

legga da standard input un valore intero che specifica il tipo di conversione da effettuare:
1: secondi (inseriti dall’utente) in ore
2: secondi inseriti dall’utente in minuti
3: minuti inseriti dall’utente in ore
4: minuti inseriti dall’utente in secondi
5: ore inserite dall’utente in secondi
6: ore inserite dall’utente in minuti
7: minuti inseriti dall’utente in giorni e ore
8: minuti inseriti dall’utente in anni e giorni

gestendo l’insertimento di un valore di scelta non compreso tra 1 e 8;

legga da standard input un valore reale da convertire;

stampi a video il valore convertito.
*/

package main

import "fmt"

func main() {

	var selezione int
	fmt.Println("Scegli la conversione:")
	fmt.Println("1) secondi -> ore")
	fmt.Println("2) secondi -> minuti")
	fmt.Println("3) minuti -> ore")
	fmt.Println("4) minuti -> secondi")
	fmt.Println("5) ore -> secondi")
	fmt.Println("6) ore -> minuti")
	fmt.Println("7) minuti -> giorni e ore")
	fmt.Println("8) minuti -> anni e giorni")
	fmt.Scan(&selezione)

	if selezione == 1 { //1) secondi -> ore
		var secondi int
		fmt.Println("Inserisci il valore da convertire: ")
		fmt.Scan(&secondi)
		ore := float64(secondi) / 3600
		fmt.Println(secondi, "secondi corrispondono a", ore, "ore")

	} else if selezione == 2 { //2) secondi -> minuti
		var secondi int
		fmt.Println("Inserisci il valore da convertire: ")
		fmt.Scan(&secondi)
		minuti := float64(secondi) / 60
		fmt.Println(secondi, "secondi corrispondono a", minuti, "minuti")

	} else if selezione == 3 { //3) minuti -> ore
		var minuti int
		fmt.Println("Inserisci il valore da convertire: ")
		fmt.Scan(&minuti)
		ore := float64(minuti) / 60
		fmt.Println(minuti, "minuti corrispondono a", ore, "ore")

	} else if selezione == 4 { //4) minuti -> secondi
		var minuti int
		fmt.Println("Inserisci il valore da convertire: ")
		fmt.Scan(&minuti)
		secondi := minuti * 60
		fmt.Println(minuti, "minuti corrispondono a", secondi, "secondi")

	} else if selezione == 5 { //5) ore -> secondi
		var ore int
		fmt.Println("Inserisci il valore da convertire: ")
		fmt.Scan(&ore)
		secondi := ore * 3600
		fmt.Println(ore, "ore corrispondono a", secondi, "secondi")

	} else if selezione == 6 { //6) ore -> minuti
		var ore int
		fmt.Println("Inserisci il valore da convertire: ")
		fmt.Scan(&ore)
		minuti := ore * 60
		fmt.Println(ore, "ore corrispondono a", minuti, "minuti")

	} else if selezione == 7 { //7) minuti -> giorni e ore
		var minuti int
		fmt.Println("Inserisci il valore da convertire: ")
		fmt.Scan(&minuti)
		giorni := minuti / 1440
		ore := float64(minuti%1440) / 60
		fmt.Println(minuti, "minuti corrispondono a", giorni, "giorni e", ore, "ore")

	} else if selezione == 8 { //8) minuti -> anni e giorni
		var minuti int
		fmt.Println("Inserisci il valore da convertire: ")
		fmt.Scan(&minuti)
		anni := minuti / (1440 * 365)
		giorni := float64(minuti%(1440*365)) / 1440
		fmt.Println(minuti, "minuti corrispondono a", anni, "anni e", giorni, "giorni")

	} else {
		fmt.Println("Scelta errata")
	}
}
