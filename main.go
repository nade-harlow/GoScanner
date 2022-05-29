package main

import "fmt"

func main() {
	err := Capture()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//RecognizeFile(filePath)

}
