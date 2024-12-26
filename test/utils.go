package test

import (
	"bufio"
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

	// read to the end of the file, should ignore CRLF/LF
	var excepted string
	reader := bufio.NewReader(fp)

	// read line by line
	for {
		var line []byte
		line, _, err = reader.ReadLine()
		if err != nil {
			break
		}
		excepted += string(line) + "\n"
	}

	return excepted
}
