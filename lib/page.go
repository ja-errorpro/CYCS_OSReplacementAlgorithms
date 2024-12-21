package lib

import (
	"time"
)

type Page struct {
	Reference byte
	Time      time.Time
	Freq      int
}

func NewPage(reference byte) *Page {
	return &Page{
		Reference: reference,
		Time:      time.Now(),
		Freq:      1,
	}
}
