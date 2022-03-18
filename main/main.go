package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func load(filePath string) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Println("cannot read file:", err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("cannot decode file:", err)
	}

	x := img.Bounds().Size().X
	y := img.Bounds().Size().Y

	fmt.Println("Successfully loaded image!")
	fmt.Printf("Image size: %d x %d", x, y)
}

func main() {
	load("/Users/j.keng/Github/go-ascii-art/ascii-pineapple.jpeg")
}
