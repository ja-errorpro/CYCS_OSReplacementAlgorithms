package test

import (
	"testing"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/cmd"
)

func TestLRUFile(t *testing.T) {

	t.Run("Test LRU algorithm with input1", func(t *testing.T) {
		_, pageFrameNumber, pageReferenceSequence := ReadInputFile("../input1_method2")
		result := cmd.LRU_Cache(pageFrameNumber, pageReferenceSequence)
		excepted := ReadExceptedFile("../out_input1_method2_bak")
		if result != excepted {
			t.Errorf("LRU algorithm failed with input1")
		}
	})

}
