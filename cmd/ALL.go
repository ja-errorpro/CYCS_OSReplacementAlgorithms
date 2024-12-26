package cmd

const (
	UNKNOWN = iota
	FIFO
	LRU
	LFU_FIFO
	MFU
	LFU_LRU
	ALL
)

func AllCache(pageFrameNumber int, pageReferenceSequence string) string {
	result := ""

	result += FIFOCache(pageFrameNumber, pageReferenceSequence) + "\n"
	result += LRU_Cache(pageFrameNumber, pageReferenceSequence) + "\n"
	result += LFU_FIFO_Cache(pageFrameNumber, pageReferenceSequence) + "\n"
	result += MFU_FIFO_Cache(pageFrameNumber, pageReferenceSequence) + "\n"
	result += LFU_LRU_Cache(pageFrameNumber, pageReferenceSequence)

	return result
}
