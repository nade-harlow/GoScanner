package main

import (
	"fmt"
	"github.com/nade-harlow/QRcode-scanner/scan"
)

func main() {
	err := scan.Capture()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//RecognizeFile(filePath)

}
