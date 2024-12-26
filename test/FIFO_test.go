package test

import (
	"testing"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/cmd"
)

// Read test data from file and test FIFO algorithm
func TestFIFOFile(t *testing.T) {

	t.Run("Test FIFO algorithm with input1", func(t *testing.T) {
		_, pageFrameNumber, pageReferenceSequence := ReadInputFile("../input1_method1")
		result := cmd.FIFOCache(pageFrameNumber, pageReferenceSequence)
		excepted := ReadExceptedFile("../out_input1_method1_bak")
		if result != excepted {
			t.Errorf("FIFO algorithm failed with input1")
		}
	})

}
