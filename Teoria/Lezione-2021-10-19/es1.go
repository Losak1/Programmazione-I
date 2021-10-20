/*
STRINGHE E CARTTERI
*/

package main

import "fmt"

func main() {
	var r rune //int32 unsigned comodo per trattare con i caratteri

	r = 0xFA
	fmt.Println(r, ":", string(r)) //se la stampo normalmente, stampa il valore numerico
	r = 'A'
	fmt.Println(r, ":", string(r)) //se la stampo convertita in stringa, stampa il carattere corrispondente
	r = 49
	fmt.Println(r, ":", string(r))

	s := "stringaè"

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i], ":", string(s[i])) //accesso ai byte della stringa (coincide con il carattere se è lungo 1 byte, cioè è ASCII)
	}

	for i, r := range s { //accesso alle rune della stringa (i caratteri)
		//r è la runa, i è la sua posizione (che posso anche non dichiarare, usando la blank variable)
		fmt.Println(i, ":", string(r))
	}

	t := "altra stringa"
	fmt.Println(s + t) //concatenazione

	s += t //altra concatenazione

	u := `Questa
	è
	una stringa multilinea
	` //è possibile creare stringhe multilinea

	fmt.Println(u)

}
