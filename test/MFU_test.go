package test

import (
	"testing"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/cmd"
)

func TestMFUFile(t *testing.T) {

	t.Run("Test MFU algorithm with input1", func(t *testing.T) {
		_, pageFrameNumber, pageReferenceSequence := ReadInputFile("../input1_method4")
		result := cmd.MFU_FIFO_Cache(pageFrameNumber, pageReferenceSequence)
		excepted := ReadExceptedFile("../out_input1_method4_bak")
		if result != excepted {
			t.Errorf("MFU algorithm failed with input1")
		}
	})

}
