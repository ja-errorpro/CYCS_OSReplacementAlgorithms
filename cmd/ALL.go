package cmd

func AllCache(pageFrameNumber int, pageReferenceSequence string) string {
	result := ""

	result += FIFOCache(pageFrameNumber, pageReferenceSequence) + "\n"
	result += LRU_Cache(pageFrameNumber, pageReferenceSequence) + "\n"

	return result
}
