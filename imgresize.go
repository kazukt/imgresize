package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	"golang.org/x/image/draw"
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

	rect := m.Bounds()
	mdst := image.NewRGBA(image.Rect(0, 0, rect.Dx()/4, rect.Dy()/4))
	draw.CatmullRom.Scale(mdst, mdst.Rect, m, rect, draw.Over, nil)

	dst, err := os.Create(dstName)
	// FIXME: update error handling.
	if err != nil {
		return err
	}
	defer dst.Close()

	// FIXME: write in various formats.
	err = png.Encode(dst, mdst)
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
