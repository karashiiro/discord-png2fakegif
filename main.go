package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/gif"
	"image/png"
	"os"
)

func main() {
	// Get the input file path
	pngFilePath := flag.String("f", "", "The PNG file to convert.")
	flag.Parse()

	if *pngFilePath == "" {
		pngFilePath = &os.Args[1]
	}

	// Load the file
	pngData, err := os.Open(*pngFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pngImage, err := png.Decode(pngData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Copy the image into a GIF canvas
	var gifBuf bytes.Buffer
	gif.Encode(&gifBuf, pngImage, &gif.Options{NumColors: 256})
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

	// Change a pixel of the first frame so that Discord interprets the GIF as animated
	gifImage.Image[0].SetColorIndex(0, 0, 0)

	// Save the GIF
	gifData, err := os.Open(*pngFilePath + ".gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gif.EncodeAll(gifData, gifImage)
}
