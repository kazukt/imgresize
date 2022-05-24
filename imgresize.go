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
	height  int
	width   int
)

func init() {
	flag.StringVar(&srcName, "src", "", "input image file")
	flag.StringVar(&dstName, "dst", "", "output image file")
	flag.IntVar(&height, "height", 30, "image height")
	flag.IntVar(&width, "width", 30, "image width")
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

	mdst := image.NewRGBA(image.Rect(0, 0, height, width))
	draw.CatmullRom.Scale(mdst, mdst.Rect, m, m.Bounds(), draw.Over, nil)

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
