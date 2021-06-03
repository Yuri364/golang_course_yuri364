package sortedmap

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var reg = regexp.MustCompile(`[^a-zA-Z0-9\s]+`)

func (s *SortedMap) SplitTextIntoWords(text string, wordLengthLimit int) {

	for _, line := range strings.Split(text, ". ") {
		if line == "" {
			continue
		}
		processedString := reg.ReplaceAllString(strings.ToLower(line), "")
		words := strings.Split(processedString, " ")

		for i, word := range words {
			word = strings.ReplaceAll(word, "\t", "")

			if utf8.RuneCountInString(word) <= wordLengthLimit {
				continue
			}

			if i == 0 || i == len(words)-1 {
				s.AddToStopLIst(word)

				continue
			}

			s.WordCounter(word)
		}
	}
}
