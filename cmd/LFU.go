package cmd

import (
	"fmt"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/lib"
)

/*
If the frequency of a page is the same, then use FIFO to replace the page.
*/
func LFU_FIFO_Cache(pageFrameNumber int, pageReference string) string {

	result := "--------------Least Frequently Used Page Replacement-----------------------\n"

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
				minIndex := FindLFUPage(lst)
				if minIndex != -1 {
					inlist[lst[minIndex].Reference] = false
					// lst[minIndex] = *page
					lst = append(lst[:minIndex], lst[minIndex+1:]...)
					lst = append(lst, *page)
				} else {
					inlist[lst[0].Reference] = false
					lst = lst[1:]
					lst = append(lst, *page)
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

func LFU_LRU_Cache(pageFrameNumber int, pageReference string) string {

	result := "--------------Least Frequently Used LRU Page Replacement-----------------------\n"

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
				minIndex := FindLFUPage(lst)
				if minIndex != -1 {
					inlist[lst[minIndex].Reference] = false
					// lst[minIndex] = *page
					lst = append(lst[:minIndex], lst[minIndex+1:]...)
					lst = append(lst, *page)
				} else {
					inlist[lst[0].Reference] = false
					lst = lst[1:]
					lst = append(lst, *page)
				}
			} else {
				lst = append(lst, *page)
			}
			inlist[pageKey] = true
		} else {
			for j, p := range lst {
				if p.Reference == pageKey {
					// move the page to the tail(head) of the list
					p.Freq++
					page = &p
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
Find the page with the least frequency.
If all the pages have the same frequency, then return -1.
*/
func FindLFUPage(lst []lib.Page) int {
	minIndex := -1
	minFreq := lst[0].Freq
	for i, p := range lst {
		if p.Freq < minFreq {
			minFreq = p.Freq
			minIndex = i
		}
	}
	return minIndex
}
