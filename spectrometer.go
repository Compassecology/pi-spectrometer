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
	fmt.Println("Saving Image From Camera")
	SaveImage()
	fmt.Println("Open Image Into Memory")
	img := OpenImage()
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			fmt.Printf("At(%v,%v) : R%v G%v B%v A%v\n", x, y, r, g, b, a)
		}
	}
}

//SaveImage from camera
//Wraper for raspistill command
//-n suppress preview
//-o output file
func SaveImage() {
	cmd := exec.Command("raspistill", "-n", "-o", "image.jpg")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}

//OpenImage and load it into memory
func OpenImage() image.Image {
	reader, err := os.Open("image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
