package main

import (
	"bytes"
	"image"
	// "image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

func main2() {
	// Generate the QR code
	qr, err := qrcode.Encode("https://example.com", qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the QR code into an image.Image
	qrImg, _, err := image.Decode(bytes.NewReader(qr))
	if err != nil {
		log.Fatal(err)
	}

	// Load your image
	imagePath := "./e10b69f0-f118-4f24-8f5a-8552004c1f33.jpg"
	img, err := loadImage(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	// Resize the image to fit inside the QR code
	img = resize.Resize(30, 0, img, resize.Lanczos3)

	// Calculate the position to center the image on the QR code
	x := (qrImg.Bounds().Dx() - img.Bounds().Dx()) / 2
	y := (qrImg.Bounds().Dy() - img.Bounds().Dy()) / 2

	// Create a new image with the QR code as the background
	result := image.NewRGBA(qrImg.Bounds())
	draw.Draw(result, qrImg.Bounds(), qrImg, image.Point{}, draw.Over)

	// Draw the image on top of the QR code
	draw.Draw(result, img.Bounds().Add(image.Point{x, y}), img, image.Point{}, draw.Over)

	// Save the result
	outputPath := "output.jpg"
	err = saveImage(result, outputPath)
	if err != nil {
		log.Fatal(err)
	}
}

// Rest of the code remains the same...

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func saveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, nil)
}
