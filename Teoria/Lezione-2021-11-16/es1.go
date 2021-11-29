/*
CREAZIONE MAPPE
*/

package main

import "fmt"

func main() {

	var m map[string]int = make(map[string]int) //lo zero value è nil

	m["mario"] = 5
	m["giovanni"] = 3
	m["andrea"] = -1

	fmt.Println(m, len(m))

	m["andrea"] = 2

	fmt.Println(m, len(m))

	v, ok := m["sebastiano"]

	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Non presente")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
