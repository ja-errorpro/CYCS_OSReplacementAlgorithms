package lib

func Contains(lst []Page, page Page) bool {
	for _, p := range lst {
		if p == page {
			return true
		}
	}
	return false
}
