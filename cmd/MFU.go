package cmd

import (
	"fmt"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/lib"
)

/*
If the frequency of a page is the same, then use FIFO to replace the page.
*/
func MFU_FIFO_Cache(pageFrameNumber int, pageReference string) string {

	result := "--------------Most Frequently Used Page Replacement -----------------------\n"

	lst := []lib.Page{}
	inlist := make(map[byte]bool) // key: page reference, value: is in the page frame

	pageFaultCount := 0
	pageReplaceCount := 0

	for i := 0; i < len(pageReference); i++ {
		pageKey := pageReference[i]
		page := lib.NewPage(pageKey)
		pageFault := !inlist[pageKey]

		result += string(pageKey) + "\t"

		if pageFault {
			pageFaultCount++
			if len(lst) == pageFrameNumber {
				pageReplaceCount++
				maxIndex := FindMFUPage(lst)
				if maxIndex != -1 { // Replace the most frequently used page
					inlist[lst[maxIndex].Reference] = false
					lst[maxIndex] = *page
				} else { // FIFO
					inlist[lst[0].Reference] = false
					lst = lst[1:]
					lst = append(lst, *page)
					// lst[0] = *page
				}
			} else {
				lst = append(lst, *page)
			}
			inlist[pageKey] = true
		} else {
			for j, p := range lst {
				if p.Reference == pageKey {
					p.Freq++
					lst[j] = p
					break
				}
			}
		}

		for j := len(lst) - 1; j >= 0; j-- {
			result += string(lst[j].Reference)
		}

		if pageFault {
			result += "\tF"
		}

		result += "\n"

	}

	statics := fmt.Sprintf("Page Fault = %d  Page Replaces = %d  Page Frames = %d\n", pageFaultCount, pageReplaceCount, pageFrameNumber)
	result += statics
	return result
}

/*
Find the page with the most frequency.
If all the pages have the same frequency, then return -1.
*/
func FindMFUPage(lst []lib.Page) int {
	maxIndex := -1
	maxFreq := lst[0].Freq
	for i, p := range lst {
		if p.Freq > maxFreq {
			maxFreq = p.Freq
			maxIndex = i
		}
	}
	return maxIndex
}
