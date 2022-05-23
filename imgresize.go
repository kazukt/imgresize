package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	src, err := os.Open("input.png")
	// FIXME: update error handling.
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	m, _, err := image.Decode(src)
	// FIXME: update error handling.
	if err != nil {
		log.Fatal(err)
	}

	dst, err := os.Create("output.png")
	// FIXME: update error handling.
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	err = png.Encode(dst, m)
	// FIXME: update error handling.
	if err != nil {
		log.Fatal(err)
	}
}
