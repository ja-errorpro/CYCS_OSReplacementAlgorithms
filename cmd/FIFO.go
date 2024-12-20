package cmd

import (
	"fmt"

	"github.com/ja-errorpro/CYCS_OSReplacementAlgorithms/lib"
)

func FIFOCache(pageFrameNumber int, pageReference string) string {

    result := "--------------FIFO-----------------------\n"

    queue := lib.NewQueue()

    pageFaultCount := 0
    pageReplaceCount := 0

    for i := 0; i < len(pageReference); i++ {
        page := pageReference[i]
        pageFault := !queue.Contains(page)
        result += string(page) + "\t"

        if queue.Len() < pageFrameNumber {
            if !queue.Contains(page) {
                queue.Push(page)
            }
        } else {
            if !queue.Contains(page) {
                queue.Pop()
                queue.Push(page)
                pageReplaceCount++
            }
        }

        lst := queue.Traverse()
        for j := len(lst) - 1; j >= 0; j-- {
            result += string(lst[j].(byte))
        }

        if pageFault {
            pageFaultCount++
            result += "\tF"
        }

        result += "\n"

    }

    statics := fmt.Sprintf("Page Fault = %d  Page Replaces = %d  Page Frames = %d\n", pageFaultCount, pageReplaceCount, pageFrameNumber)
    result += statics

    return result
}
