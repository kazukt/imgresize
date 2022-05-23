package main

import (
	"flag"
	"image"
	"image/png"
	"log"
	"os"
)

var (
	srcName string
	dstName string
)

func init() {
	flag.StringVar(&srcName, "src", "", "input image file")
	flag.StringVar(&dstName, "dst", "", "output image file")
}

func main() {
	flag.Parse()
	if srcName == "" {
		log.Fatal("input file is required")
	}
	if dstName == "" {
		log.Fatal("output file is required")
	}

	src, err := os.Open(srcName)
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

	dst, err := os.Create(dstName)
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
