package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
}

func ReadFile(inputFileName string) (string, int, string) {
	fp, err := os.Open(inputFileName)
	if err != nil {
		fp, err = os.Open(inputFileName + ".txt")
		if err != nil {
			log.Fatal(err)
			return "", 0, ""
		}
	}
	defer fp.Close()

	var method string
	var pageFrameNumber int

	fmt.Fscanf(fp, "%s %d\n", &method, &pageFrameNumber)

	var pageReferenceSequence string
	fmt.Fscanf(fp, "%s\n", &pageReferenceSequence)

	return method, pageFrameNumber, pageReferenceSequence
}

func main() {
	var inputFileName string
	fmt.Print("Please enter File Name (eg. input1 、 input1.txt) : ")
	fmt.Scanln(&inputFileName)

	/*method, pageFrameNumber, pageReferenceSequence := ReadFile(inputFileName)

	switch method {
	case "1":
	case "2":
	case "3":
	case "4":
	case "5":
	case "6":
	default:
		log.Fatal("Invalid method")
	}*/

}
