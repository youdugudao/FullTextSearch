package search

import (
	"github.com/huichen/sego"
	"full_text_search/config"
	"strings"
)

var Seg sego.Segmenter

func InitSeg() {
	Seg.LoadDictionary("dictionary.txt")
}

func FilterWords(s string) bool {
	var sNew []rune
	for _, v := range s {
		if _, ok := config.FilterWords[string(v)]; ok {
			continue
		}
		if strings.TrimSpace(string(v)) == "" {
			continue
		}
		sNew = append(sNew, v)
	}
	if len(sNew) > 0 {
		return true
	} else {
		return false
	}
}

func Segment(text []byte) (words []string) {
	segments := Seg.Segment(text)
	wordsOld := sego.SegmentsToSlice(segments, true)
	for _, v := range wordsOld {
		if !FilterWords(v) {
			continue
		}
		words = append(words, v)
	}
	return
}
