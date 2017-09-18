package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Raspberry PI Spectrometer")

	SaveImage()
	img := OpenImage()
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			fmt.Printf("At(%v,%v) : R%v G%v B%v A%v", x, y, r, g, b, a)
		}
	}
}

//SaveImage from camera
func SaveImage() {
	cmd := exec.Command("raspistill", "-w", "-o", "image.jpg")
	_, err := cmd.StdoutPipe()
	LogFatal(err)
	err = cmd.Start()
	LogFatal(err)
	cmd.Wait()
}

//OpenImage and load it into memory
func OpenImage() image.Image {
	reader, err := os.Open("image.jpg")
	LogFatal(err)
	defer reader.Close()
	img, _, err := image.Decode(reader)
	LogFatal(err)
	return img
}

//LogFatal quit program if we get an error
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
