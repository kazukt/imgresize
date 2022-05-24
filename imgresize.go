package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/png"
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

func run() error {
	flag.Parse()
	if srcName == "" {
		return errors.New("input file is required")
	}
	if dstName == "" {
		return errors.New("output file is required")
	}

	src, err := os.Open(srcName)
	// FIXME: update error handling.
	if err != nil {
		return err
	}
	defer src.Close()

	m, _, err := image.Decode(src)
	// FIXME: update error handling.
	if err != nil {
		return err
	}

	dst, err := os.Create(dstName)
	// FIXME: update error handling.
	if err != nil {
		return err
	}
	defer dst.Close()

	err = png.Encode(dst, m)
	// FIXME: update error handling.
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
