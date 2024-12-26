package test

import (
	"fmt"
	"log"
	"os"
)

func ReadInputFile(inputFileName string) (int, int, string) {
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

func ReadExceptedFile(fileName string) string {
	fp, err := os.Open(fileName)
	if err != nil {
		fp, err = os.Open(fileName + ".txt")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	defer fp.Close()

	// should read to the end of the file
	buf := make([]byte, 10240)
	n, err := fp.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf[:n])
}
