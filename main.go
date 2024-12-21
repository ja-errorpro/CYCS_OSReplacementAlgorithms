package main

import (
	"fmt"
	"log"
	"os"

	methods "github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/cmd"
	_ "github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/init"
)

func ReadFile(inputFileName string) (int, int, string) {
	fp, err := os.Open(inputFileName)
	if err != nil {
		fp, err = os.Open(inputFileName + ".txt")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	defer fp.Close()

	var method int
	var pageFrameNumber int

	fmt.Fscanf(fp, "%d %d\n", &method, &pageFrameNumber)

	var pageReferenceSequence string
	fmt.Fscanf(fp, "%s\n", &pageReferenceSequence)

	return method, pageFrameNumber, pageReferenceSequence
}

func WriteFile(outputFileName string, buf string) {
	fp, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer fp.Close()

	fmt.Fprint(fp, buf)
}

func main() {
	var inputFileName string
	fmt.Print("Please enter File Name (eg. input1 、 input1.txt) : ")
	fmt.Scanln(&inputFileName)

	method, pageFrameNumber, pageReferenceSequence := ReadFile(inputFileName)

	var outputFileName string = "out_" + inputFileName

	if outputFileName[len(outputFileName)-4:] != ".txt" {
		outputFileName += ".txt"
	}

	var result string

	switch method {
	case methods.FIFO:
		result = methods.FIFOCache(pageFrameNumber, pageReferenceSequence)
	case methods.LRU:
		result = methods.LRU_Cache(pageFrameNumber, pageReferenceSequence)
	case methods.LFU_FIFO:
	case methods.MFU:
	case methods.LFU_LRU:
	case methods.ALL:
		result = methods.AllCache(pageFrameNumber, pageReferenceSequence)
	default:
		log.Fatal("Invalid method")
	}

	WriteFile(outputFileName, result)
	fmt.Print(result)

}
