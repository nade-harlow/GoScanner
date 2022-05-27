package main

import (
	"bytes"
	"fmt"
	"github.com/liyue201/goqr"
	"gocv.io/x/gocv"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"time"
)

var (
	fileName = fmt.Sprint(rand.Intn(100000)) + ".png"
	filePath = "testdata/" + fileName
	device   = "0"
)

func main() {
	capture()
	recognizeFile(filePath)

}

// capture opens the webcam and captures a frame
func capture() {
	sig := make(chan bool)
	go func(deviceID string, saveFile string, sig chan bool) {
		fmt.Println("openning device: ", device)

		webcam, err := gocv.OpenVideoCapture(device)
		if err != nil {
			fmt.Printf("Error opening video capture device: %v\n", device)
			return
		}
		time.Sleep(time.Second * 3)
		fmt.Printf("start reading device: %v\n", device)
		defer webcam.Close()

		// image matrix to hold captured image from webcam in memory for processing and saving to file
		img := gocv.NewMat()
		defer img.Close()
		fmt.Printf("Capturing image\n")

		// read image from webcam into img matrix. ok = true if read successful and false if not
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Error reading from device %v\n", device)
			return
		}
		if img.Empty() { // if img is empty, then webcam is not working
			fmt.Printf("no image on device %v\n", device)
		}
		fmt.Printf("Saving image\n")
		// save image to file ex.jpeg in current directory
		gocv.IMEncode(".png", img)
		gocv.IMWrite(filePath, img)
		fmt.Printf("\nFinished\n")

		sig <- true
	}(device, fileName, sig)

	<-sig
}

// recognizeFile recognize qr code from file
func recognizeFile(path string) {
	fmt.Printf("recognize file: %v\n", path)
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	//go os.Remove(path)
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
		return
	}
	for _, qrCode := range qrCodes {
		fmt.Printf("QRcode payload: %s\n", qrCode.Payload)
	}

	return
}
