/*
TESTO DEL PROBLEMA
*/

package main

import (
	"image"
	"image/png"
	"os"
)

func main() {

	i := image.NewRGBA(image.Rect(0, 0, 600, 600))

	png.Encode(os.Stdout, i)
}
