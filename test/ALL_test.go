package test

import (
	"testing"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/cmd"
)

func TestALLFile(t *testing.T) {

	t.Run("Test input1", func(t *testing.T) {
		_, pageFrameNumber, pageReferenceSequence := ReadInputFile("../input1_method6")
		result := cmd.AllCache(pageFrameNumber, pageReferenceSequence)
		excepted := ReadExceptedFile("../out_input1_method6_bak")
		if result != excepted {
			t.Errorf("Test failed with input1_method6")
		}
	})

	t.Run("Test input2", func(t *testing.T) {
		_, pageFrameNumber, pageReferenceSequence := ReadInputFile("../input2")
		result := cmd.AllCache(pageFrameNumber, pageReferenceSequence)
		excepted := ReadExceptedFile("../out_input2")
		if result != excepted {
			t.Errorf("Test failed with input2")
		}
	})

}
