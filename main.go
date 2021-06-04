package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
)

type transparencyQuantizer struct{}

func (tq *transparencyQuantizer) Quantize(p color.Palette, m image.Image) color.Palette {
	customPalette := palette.Plan9
	customPalette[0] = color.Transparent
	return customPalette
}

func main() {
	// Get the input file path
	if len(os.Args) == 0 {
		fmt.Println("no input file specified")
		os.Exit(1)
	}

	pngFilePath := os.Args[1]

	// Load the file
	pngData, err := os.Open(pngFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer pngData.Close()

	pngImage, err := png.Decode(pngData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Copy the image into a GIF canvas
	var gifBuf bytes.Buffer
	gif.Encode(&gifBuf, pngImage, &gif.Options{
		NumColors: 256,
		Quantizer: &transparencyQuantizer{},
		Drawer:    draw.Src,
	})

	gifImage, err := gif.DecodeAll(&gifBuf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Animate GIF
	dupeFrame := *gifImage.Image[0]
	gifImage.Delay = []int{0, 0}
	gifImage.Image = append(gifImage.Image, &dupeFrame)
	gifImage.LoopCount = -1
	gifImage.Disposal = nil

	// Change a pixel of the first frame so that Discord interprets the GIF as animated
	gifImage.Image[0].SetColorIndex(0, 0, 0)

	// Save the GIF
	gifData, err := os.Create(pngFilePath + ".gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer gifData.Close()

	gif.EncodeAll(gifData, gifImage)
}
