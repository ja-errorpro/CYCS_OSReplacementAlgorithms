package cmd

import (
	"fmt"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/lib"
)

/*
This function use stack implementation to simulate the LRU algorithm.
Time Complexity: O(1)
*/
func LRU_Cache(pageFrameNumber int, pageReference string) string {

	result := "--------------LRU-----------------------\n"

	lst := []lib.Page{}
	hashMap := make(map[byte]bool) // key: page reference, value: is in the page frame

	pageFaultCount := 0
	pageReplaceCount := 0

	for i := 0; i < len(pageReference); i++ {
		pageKey := pageReference[i]
		page := lib.NewPage(pageKey)
		pageFault := !hashMap[pageKey]
		result += string(pageKey) + "\t"

		if pageFault {
			pageFaultCount++
			if len(lst) == pageFrameNumber {
				pageReplaceCount++
				hashMap[lst[0].Reference] = false
				lst = lst[1:]
			}
			lst = append(lst, *page)
			hashMap[pageKey] = true
		} else {
			for j, p := range lst {
				if p.Reference == pageKey {
					// move the page to the tail(head) of the list
					lst = append(lst[:j], lst[j+1:]...)
					lst = append(lst, *page)
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
This function use counter implementation to simulate the LRU algorithm.
Time Complexity: O(n)
*/
func LRU_Cache_Counter(pageFrameNumber int, pageReference string) string {

	result := "--------------LRU-----------------------\n"

	lst := []lib.Page{}

	pageFaultCount := 0
	pageReplaceCount := 0

	for i := 0; i < len(pageReference); i++ {
		pageKey := pageReference[i]
		page := lib.NewPage(pageKey)
		pageFault := !lib.Contains(lst, *page)
		result += string(pageKey) + "\t"

		if pageFault {
			pageFaultCount++
			if len(lst) == pageFrameNumber {
				pageReplaceCount++
				replaceTarget := 0

				for j := 1; j < len(lst); j++ {
					if lst[j].Time.Before(lst[replaceTarget].Time) {
						replaceTarget = j
					}
				}

				lst[replaceTarget] = *page
			} else {
				lst = append(lst, *page)
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
