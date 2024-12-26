package test

import (
	"testing"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/cmd"
)

func TestLFUFIFOFile(t *testing.T) {

	t.Run("Test LFU_FIFO algorithm with input1", func(t *testing.T) {
		_, pageFrameNumber, pageReferenceSequence := ReadInputFile("../input1_method3")
		result := cmd.LFU_FIFO_Cache(pageFrameNumber, pageReferenceSequence)
		excepted := ReadExceptedFile("../out_input1_method3_bak")
		if result != excepted {
			t.Errorf("LFU_FIFO algorithm failed with input1")
		}
	})

}

func TestLFULRUFile(t *testing.T) {
	t.Run("Test LFU_LRU algorithm with input1", func(t *testing.T) {
		_, pageFrameNumber, pageReferenceSequence := ReadInputFile("../input1_method5")
		result := cmd.LFU_LRU_Cache(pageFrameNumber, pageReferenceSequence)
		excepted := ReadExceptedFile("../out_input1_method5_bak")
		if result != excepted {
			t.Errorf("LFU_LRU algorithm failed with input1")
		}
	})
}
