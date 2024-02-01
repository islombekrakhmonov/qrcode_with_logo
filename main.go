package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"os/exec"

	"github.com/divan/qrlogo"
	"github.com/nfnt/resize"
	qr "github.com/skip2/go-qrcode"
)

func main() {

	logoImage, err := getImageFromFilePath("./ppg_logo.jpg")
	if err != nil {
		log.Fatal(err)
	}

	

	resizedLogo := resize.Resize(300, 300, logoImage, resize.Lanczos3)

	encoder := qrlogo.Encoder{
		AlphaThreshold: 2000,
		GreyThreshold:  30,
		QRLevel:        qr.Highest,
	}

	qrCode, err := encoder.Encode("PPLLNG305UFUE7", resizedLogo, 1000)
	if err != nil {
		log.Fatal(err)
	}

	qrImage, _, err := image.Decode(qrCode)
	if err != nil {
		log.Fatal(err)
	}

	qrImageResized := resize.Resize(256, 256, qrImage, resize.Lanczos3)

	outputFilename := "qrcode_with_logo.jpg"

	err = saveImages(qrImageResized, "qrcode_with_logo.jpg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("QR code generated with logo successfully")

	err = openFile(outputFilename)
	if err != nil {
		log.Fatal(err)
	}
}

func openFile(filename string) error {
	// cmd := exec.Command("xdg-open", filename) // for Linux
	cmd := exec.Command("open", filename)    // for macOS
	// cmd := exec.Command("start", filename)   // for Windows

	err := cmd.Run()
	return err
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func saveImages(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}

	return nil
}


func getOriginalLogoDimensions(logoImage image.Image) (int, int) {
	// Implement your logic to get the original dimensions of the logo
	// Replace this with your actual implementation.
	return logoImage.Bounds().Dx(), logoImage.Bounds().Dy()
}

func calculateResizedDimensions(originalWidth, originalHeight int) (int, int) {
	// Implement your logic to calculate the resized dimensions
	// You may adjust this logic based on your requirements.
	// Here, we maintain the original aspect ratio and resize to a maximum of 300 pixels in either dimension.
	const maxDimension = 300

	var resizedWidth, resizedHeight int

	if originalWidth > originalHeight {
		resizedWidth = maxDimension
		resizedHeight = int(float64(originalHeight) * (float64(maxDimension) / float64(originalWidth)))
	} else {
		resizedHeight = maxDimension
		resizedWidth = int(float64(originalWidth) * (float64(maxDimension) / float64(originalHeight)))
	}

	return resizedWidth, resizedHeight
}

func calculateQRCodeSize(resizedWidth, resizedHeight int) int {
	// Implement your logic to calculate the QR code size based on the resized logo dimensions
	// You may adjust this logic based on your requirements.
	// Here, we use a percentage of the resized logo's width as the QR code size.
	const percentageOfWidthForQRCode = 50

	qrCodeSize := (percentageOfWidthForQRCode * resizedWidth) / 100
	return qrCodeSize
}