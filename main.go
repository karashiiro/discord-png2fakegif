package main

import (
	"flag"
	"log"
)

func main() {
	pngFilePath := flag.String("f", "", "The PNG file to convert.")
	flag.Parse()

	log.Println(*pngFilePath)
}
